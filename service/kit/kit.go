package kit

import (
	"context"
	"e-commerce/service/app"
	"e-commerce/service/infra/log"
	"e-commerce/service/infra/routine"
	"e-commerce/service/kit/endpoint"
	"e-commerce/service/kit/transport"
	pb "e-commerce/service/kit/transport/grpc"
	"flag"
	"fmt"
	klog "github.com/go-kit/kit/log"
	kitTransport "github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/xxl-job/xxl-job-executor-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// definition server after
var traceOptions = otgrpc.SpanDecorator(func(ctx context.Context, span opentracing.Span, method string, req, resp interface{}, grpcError error) {
	span.SetTag("monitor.api", method)
})

func startHttpServer(svcEndpoint endpoint.Set) {
	var (
		httpAddr = flag.String("http.addr", ":80", "HTTP listen address")
	)
	flag.Parse()

	var logger klog.Logger
	{
		logger = klog.NewLogfmtLogger(os.Stderr)
		logger = klog.With(logger, "ts", klog.DefaultTimestampUTC)
		logger = klog.With(logger, "caller", klog.DefaultCaller)
	}

	var h http.Handler
	{
		options := []httptransport.ServerOption{
			httptransport.ServerErrorHandler(kitTransport.NewLogErrorHandler(klog.With(logger, "component", "HTTP"))),
			httptransport.ServerErrorEncoder(transport.EncodeProjectError),
		}

		// handler
		h = transport.MakeHTTPHandler(svcEndpoint, options)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
	return
}

var Exec xxl.Executor

func runGrpcServer(svcEndpoint endpoint.Set) {
	host := "0.0.0.127"

	log.Info("*** Starting grpc server ", host)
	ctx := routine.GetChildContext()

	grpcServer := transport.NewGrpcServer(svcEndpoint)
	grpcListener, err := net.Listen("tcp", host)
	if err != nil {
		log.Errorf("transport gRPC during Listen err: %v", err)
	}
	maxMsgSize := 1024 * 1024 * 100
	baseServer := grpc.NewServer(
		// grpc.UnaryInterceptor(kitgrpc.Interceptor),
		grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer(), traceOptions)),
		grpc.MaxRecvMsgSize(maxMsgSize),
		grpc.MaxSendMsgSize(maxMsgSize))
	reflection.Register(baseServer)
	pb.RegisterRetailServiceServer(baseServer, grpcServer)

	healthSvc := health.NewServer()
	healthpb.RegisterHealthServer(baseServer, healthSvc)

	exitGrpcServer := func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("close grpc server")
				baseServer.Stop()
				return
			}
		}
	}
	routine.RunAsDaemon(exitGrpcServer)

	log.Info("Grpc server started")
	baseServer.Serve(grpcListener)
}

func registerCronjob(svcEndpoint endpoint.Set) {

}

func StartMainCenterGrpcService() {
	mainService := app.NewService()
	svcEndpoint := endpoint.New(mainService)

	routine.RunAsDaemon(func() {
		startHttpServer(svcEndpoint)
	})
	//routine.RunAsDaemon(func() {
	//	runGrpcServer(svcEndpoint)
	//})
	//routine.RunAsDaemon(func() {
	//	registerCronjob(svcEndpoint)
	//})
}
