package trace

import (
	"context"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	openLog "github.com/opentracing/opentracing-go/log"
	zipkinopentracing "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"runtime"
	"strings"
)

// metadata 读写
type MDReaderWriter struct {
	metadata.MD
}

// 为了 opentracing.TextMapReader ，参考 opentracing 代码
func (c MDReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range c.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// 为了 opentracing.TextMapWriter，参考 opentracing 代码
func (c MDReaderWriter) Set(key, val string) {
	key = strings.ToLower(key)
	c.MD[key] = append(c.MD[key], val)
}

func MakeTraceServerInterceptor(tracer opentracing.Tracer, service string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		md, succ := metadata.FromIncomingContext(ctx)
		if !succ {
			md = metadata.New(nil)
		}
		grpclog.Errorf("extract from metadata md: %v", md)

		// 提取 spanContext
		spanContext, err := tracer.Extract(opentracing.TextMap, MDReaderWriter{md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			grpclog.Errorf("extract from metadata err: %v", err)
		} else {
			span := tracer.StartSpan(
				info.FullMethod,
				ext.RPCServerOption(spanContext),
				opentracing.Tag{Key: string(ext.Component), Value: "grpc"},
				ext.SpanKindRPCServer,
			)
			span.SetTag("monitor.api", service)
			defer span.Finish()
			ctx = opentracing.ContextWithSpan(ctx, span)
		}

		return kitgrpc.Interceptor(ctx, req, info, handler)
	}
}

func MakeTraceClientInterceptor(tracer opentracing.Tracer, service string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context,
		method string,
		req,
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {
		// 创建 rootSpan
		var rootCtx opentracing.SpanContext

		rootSpan := opentracing.SpanFromContext(ctx)
		if rootSpan != nil {
			rootCtx = rootSpan.Context()
		}

		span := tracer.StartSpan(
			method,
			opentracing.ChildOf(rootCtx),
			opentracing.Tag{"trace", "client"},
			ext.SpanKindRPCClient,
		)
		span.SetTag("monitor.api", service)
		defer span.Finish()

		md, succ := metadata.FromOutgoingContext(ctx)
		if !succ {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		mdWriter := MDReaderWriter{md}

		// 注入 spanContext
		err := tracer.Inject(span.Context(), opentracing.TextMap, mdWriter)

		if err != nil {
			span.LogFields(openLog.String("inject error", err.Error()))
		}

		// new ctx ，并调用后续操作
		newCtx := metadata.NewOutgoingContext(ctx, md)
		err = invoker(newCtx, method, req, reply, cc, opts...)
		if err != nil {
			span.LogFields(openLog.String("call error", err.Error()))
		}
		return err
	}
}

func NewZipKinTracer(serviceName string) opentracing.Tracer {
	endPointUrl := os.Getenv("ZIPKIN_BASE_URL")
	//zipkinreporter.Timeout 上报链路日志超时时间（http）
	//zipkinreporter.BatchSize 每次推送数量
	//zipkinreporter.BatchInterval 批量推送周期
	//zipkinreporter.MaxBacklog 链路日志缓冲区大小，最大1000，超过1000会被丢弃
	reporter := zipkinreporter.NewReporter(endPointUrl)

	//create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(serviceName, "localhost:0")
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}
	//zipkin.NewModuloSampler
	// initialize our tracer
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}

	// use zipkin-go-opentracing to wrap our tracer
	tracer := zipkinopentracing.Wrap(nativeTracer)

	opentracing.SetGlobalTracer(tracer)

	return tracer
}

// SpanWithCtx 生成上下文span
func SpanWithCtx(ctx context.Context) (opentracing.Span, context.Context) {
	//开始链路追踪
	pc, _, _, _ := runtime.Caller(1)
	spanName := ""
	if pc > 0 {
		spanName = spanName + "/" + runtime.FuncForPC(pc).Name()
	}
	return opentracing.StartSpanFromContext(ctx, spanName)
}
