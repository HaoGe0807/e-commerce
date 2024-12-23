package http

import (
	"context"
	"e-commerce/service/kit/endpoint"
	"encoding/json"
	"fmt"
	kitEndpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func CreateProductServer(endpoint kitEndpoint.Endpoint, options ...httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		decodeCreateProductRequest,
		encodeResp,
		options...,
	)
}
func decodeCreateProductRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.CreateProductReq
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	err = FormatParam(ctx, req)
	if err != nil {
		//return nil, &errors.Err{
		//	Code: http.StatusBadRequest,
		//	Msg:  err.Error(),
		//	Data: nil,
		//}
		return nil, err
	}
	return req, nil
}

func UpdateProductServer(endpoint kitEndpoint.Endpoint, options ...httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		decodeUpdateProductRequest,
		encodeResp,
		options...,
	)
}
func decodeUpdateProductRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.UpdateProductReq
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	err = FormatParam(ctx, req)
	if err != nil {
		//return nil, &errors.Err{
		//	Code: http.StatusBadRequest,
		//	Msg:  err.Error(),
		//	Data: nil,
		//}
		return nil, err
	}
	return req, nil
}

func DeleteProductServer(endpoint kitEndpoint.Endpoint, options ...httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		decodeDeleteProductRequest,
		encodeResp,
		options...,
	)
}
func decodeDeleteProductRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.DeleteProductReq
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	err = FormatParam(ctx, req)
	if err != nil {
		//return nil, &errors.Err{
		//	Code: http.StatusBadRequest,
		//	Msg:  err.Error(),
		//	Data: nil,
		//}
		return nil, err
	}
	return req, nil
}

func QueryProductServer(endpoint kitEndpoint.Endpoint, options ...httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		decodeQueryProductRequest,
		encodeResp,
		options...,
	)
}
func decodeQueryProductRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.QueryProductReq
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	err = FormatParam(ctx, req)
	if err != nil {
		//return nil, &errors.Err{
		//	Code: http.StatusBadRequest,
		//	Msg:  err.Error(),
		//	Data: nil,
		//}
		return nil, err
	}
	return req, nil
}

func QueryProductListServer(endpoint kitEndpoint.Endpoint, options ...httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		decodeQueryProductListRequest,
		encodeResp,
		options...,
	)
}
func decodeQueryProductListRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.QueryProductListReq
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	err = FormatParam(ctx, req)
	if err != nil {
		//return nil, &errors.Err{
		//	Code: http.StatusBadRequest,
		//	Msg:  err.Error(),
		//	Data: nil,
		//}
		return nil, err
	}
	return req, nil
}

func CreateCategoryServer(endpoint kitEndpoint.Endpoint, options ...httptransport.ServerOption) *httptransport.Server {
	fmt.Print("test")
	return httptransport.NewServer(
		endpoint,
		decodeCreateCategoryRequest,
		encodeResp,
		options...,
	)
}
func decodeCreateCategoryRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.CreateCategoryReq

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	err = FormatParam(ctx, req)
	if err != nil {
		//return nil, &errors.Err{
		//	Code: http.StatusBadRequest,
		//	Msg:  err.Error(),
		//	Data: nil,
		//}
		return nil, err
	}
	return req, nil
}

func UpdateCategoryServer(endpoint kitEndpoint.Endpoint, options ...httptransport.ServerOption) *httptransport.Server {

	return httptransport.NewServer(
		endpoint,
		decodeUpdateCategoryRequest,
		encodeResp,
		options...,
	)
}
func decodeUpdateCategoryRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.UpdateCategoryReq

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	err = FormatParam(ctx, req)
	if err != nil {
		//return nil, &errors.Err{
		//	Code: http.StatusBadRequest,
		//	Msg:  err.Error(),
		//	Data: nil,
		//}
		return nil, err
	}
	return req, nil
}

func DeleteCategoryServer(endpoint kitEndpoint.Endpoint, options ...httptransport.ServerOption) *httptransport.Server {

	return httptransport.NewServer(
		endpoint,
		decodeDeleteCategoryRequest,
		encodeResp,
		options...,
	)
}
func decodeDeleteCategoryRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.DeleteCategoryReq

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	err = FormatParam(ctx, req)
	if err != nil {
		//return nil, &errors.Err{
		//	Code: http.StatusBadRequest,
		//	Msg:  err.Error(),
		//	Data: nil,
		//}
		return nil, err
	}
	return req, nil
}

func QueryCategoryServer(endpoint kitEndpoint.Endpoint, options ...httptransport.ServerOption) *httptransport.Server {

	return httptransport.NewServer(
		endpoint,
		decodeQueryCategoryRequest,
		encodeResp,
		options...,
	)
}
func decodeQueryCategoryRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.QueryCategoryReq

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	err = FormatParam(ctx, req)
	if err != nil {
		//return nil, &errors.Err{
		//	Code: http.StatusBadRequest,
		//	Msg:  err.Error(),
		//	Data: nil,
		//}
		return nil, err
	}
	return req, nil
}

func QueryCategoryListServer(endpoint kitEndpoint.Endpoint, options ...httptransport.ServerOption) *httptransport.Server {

	return httptransport.NewServer(
		endpoint,
		decodeQueryCategoryListRequest,
		encodeResp,
		options...,
	)
}
func decodeQueryCategoryListRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.QueryCategoryListReq

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	err = FormatParam(ctx, req)
	if err != nil {
		//return nil, &errors.Err{
		//	Code: http.StatusBadRequest,
		//	Msg:  err.Error(),
		//	Data: nil,
		//}
		return nil, err
	}
	return req, nil
}
