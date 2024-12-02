package transport

import (
	"bytes"
	"context"
	errors "e-commerce/service/infra/errors"
	"e-commerce/service/kit/endpoint"
	kitHttp "e-commerce/service/kit/transport/http"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/pprof"
	"strings"
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
// func MakeHTTPHandler(svcEndpoint endpoint.Set, options []httptransport.ServerOption,openTracer opentracing.Tracer) http.Handler {
func MakeHTTPHandler(svcEndpoint endpoint.Set, options []httptransport.ServerOption) http.Handler {
	traceMiddleware := func(h http.Handler) http.Handler {
		return nethttp.Middleware(
			opentracing.GlobalTracer(),
			h,
			nethttp.OperationNameFunc(func(r *http.Request) string {
				return r.URL.EscapedPath()
			}),
			nethttp.MWSpanObserver(func(sp opentracing.Span, r *http.Request) {
				sp.SetTag("monitor.api", r.URL.EscapedPath())
			}),
		)
	}
	r := mux.NewRouter()

	// pprof
	{
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.HandleFunc("/debug/pprof/trace", pprof.Trace)
		r.PathPrefix("/debug/pprof/").HandlerFunc(pprof.Index)
	}
	// prometheus
	{
		r.Handle("/monitor/prometheus", promhttp.Handler())
	}

	// use tracing middleware
	r.Use(traceMiddleware)
	// check
	r.Methods("GET").Path("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	// product
	productRouter := r.PathPrefix("/api/e-commerce/product").Subrouter()
	{
		//productRouter.Use(formDataToJsonMiddleware)
		productRouter.Methods("POST").Path("/createProduct").Handler(kitHttp.CreateProductServer(svcEndpoint.CreateProductEP, options...))
		productRouter.Methods("POST").Path("/updateProduct").Handler(kitHttp.UpdateProductServer(svcEndpoint.UpdateProductEP, options...))
		productRouter.Methods("POST").Path("/deleteProduct").Handler(kitHttp.DeleteProductServer(svcEndpoint.DeleteProductEP, options...))
		productRouter.Methods("POST").Path("/queryProduct").Handler(kitHttp.QueryProductServer(svcEndpoint.QueryProductEP, options...))
		productRouter.Methods("POST").Path("/queryProductList").Handler(kitHttp.QueryProductListServer(svcEndpoint.QueryProductListEP, options...))

		//category
		//productRouter.Methods("POST").Path("/category/createCategory").Handler(kitHttp.CreateCategoryServer(svcEndpoint.CreateCategoryEP, options...))
		//productRouter.Methods("POST").Path("/category/updateCategory").Handler(kitHttp.UpdateCategoryServer(svcEndpoint.UpdateCategoryEP, options...))
		//productRouter.Methods("POST").Path("/category/deleteCategory").Handler(kitHttp.DeleteCategoryServer(svcEndpoint.DeleteCategoryEP, options...))
		//productRouter.Methods("POST").Path("/category/queryCategory").Handler(kitHttp.QueryCategoryServer(svcEndpoint.QueryCategoryEP, options...))
		//productRouter.Methods("POST").Path("/category/queryCategoryList").Handler(kitHttp.QueryCategoryListServer(svcEndpoint.QueryCategoryListEP, options...))
	}

	return r
}

/**
 * @Author YBH
 * @Description // 解析业务错误
 * @Date 3:49 下午 2022/9/26
 * @Param
 * @return
 **/
func EncodeProjectError(ctx context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain;applicatiion/json;charset=utf-8", []byte(err.Error())
	w.Header().Set("Content-Type", contentType)
	if myerr, ok := err.(*errors.Err); ok {
		w.WriteHeader(http.StatusOK)
		retErr := &kitHttp.Response{
			Code: int(myerr.Code),
			Msg:  myerr.Msg,
			Data: myerr.Data,
		}
		errS, _ := json.Marshal(retErr)

		w.Write(errS)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
	}

}

func formDataToJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			if strings.Contains(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
				err := r.ParseForm()
				if err != nil {
					return
				}
			} else {
				err := r.ParseMultipartForm(10000)
				if err != nil {
					return
				}
			}
			params := r.Form.Get("params")
			buff := []byte(params)
			r.Header.Set("Content-Type", "application/json")
			r.Body = ioutil.NopCloser(bytes.NewBuffer(buff))
		} else {
			obj := make(map[string]interface{})
			buf, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
			err = json.Unmarshal(buf, &obj)
			if err != nil {
				return
			}
			if obj["params"] != nil {
				if params, ok := obj["params"].(string); ok {
					paramsBuf := []byte(params)
					r.Body = ioutil.NopCloser(bytes.NewBuffer(paramsBuf))
				}
			} else {
				r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
			}
		}
		next.ServeHTTP(w, r)
	})
}
