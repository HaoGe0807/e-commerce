package endpoint

import (
	"context"
	"e-commerce/service/app"
	"e-commerce/service/infra/utils"
	goEndpoint "github.com/go-kit/kit/endpoint"
)

type Set struct {
	CreateProductEP    goEndpoint.Endpoint
	UpdateProductEP    goEndpoint.Endpoint
	DeleteProductEP    goEndpoint.Endpoint
	QueryProductEP     goEndpoint.Endpoint
	QueryProductListEP goEndpoint.Endpoint
}

func New(service app.ECommerceService) Set {
	createProductEP := NewCreateProductEP(service)
	createProductEP = utils.LoggingMiddleware()(createProductEP)
	createProductEP = utils.ParameterCheckMiddleware()(createProductEP)

	updateProductEP := NewUpdateProductEP(service)
	updateProductEP = utils.LoggingMiddleware()(updateProductEP)
	updateProductEP = utils.ParameterCheckMiddleware()(updateProductEP)

	deleteProductEP := NewDeleteProductEP(service)
	deleteProductEP = utils.LoggingMiddleware()(deleteProductEP)
	deleteProductEP = utils.ParameterCheckMiddleware()(deleteProductEP)

	queryProductEP := NewQueryProductEP(service)
	queryProductEP = utils.LoggingMiddleware()(queryProductEP)
	queryProductEP = utils.ParameterCheckMiddleware()(queryProductEP)

	queryProductListEP := NewQueryProductListEP(service)
	queryProductListEP = utils.LoggingMiddleware()(queryProductListEP)
	queryProductListEP = utils.ParameterCheckMiddleware()(queryProductListEP)
	return Set{
		CreateProductEP:    createProductEP,
		UpdateProductEP:    updateProductEP,
		DeleteProductEP:    deleteProductEP,
		QueryProductEP:     queryProductEP,
		QueryProductListEP: queryProductListEP,
	}
}

func NewCreateProductEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateProductReq)

		err = service.CreateProduct(ctx, req.ProductName, req.CategoryId, req.Skus, req.Unit, req.MnemonicCode, req.Status, req.StoreId, req.CustomizationList, req.IngredientList, req.Icon, req.PriceMethod, req.Shape, req.ShapeColor, req.FirstDisplay, req.ProductSpecifications, req.ProductType)
		if err != nil {
			return nil, err
		}
		resp := &CreateProductResp{}
		return resp, nil
	}
}

func NewUpdateProductEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateProductReq)

		err = service.UpdateProduct(ctx, req.SpuId, req.ProductName, req.CategoryId, req.Skus, req.Unit, req.MnemonicCode, req.Status, req.StoreId, req.CustomizationList, req.IngredientList, req.Icon, req.PriceMethod, req.Shape, req.ShapeColor, req.FirstDisplay, req.ProductSpecifications, req.ProductType)
		if err != nil {
			return nil, err
		}
		resp := &UpdateProductResp{}
		return resp, nil
	}
}

func NewDeleteProductEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteProductReq)

		err = service.DeleteProduct(ctx, req.StoreId, req.SpuId)
		if err != nil {
			return nil, err
		}
		resp := &DeleteProductResp{}
		return resp, nil
	}
}

func NewQueryProductEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(QueryProductReq)

		product, err := service.QueryProduct(ctx, req.StoreId, req.SpuId)
		if err != nil {
			return nil, err
		}

		return product, nil
	}
}

func NewQueryProductListEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(QueryProductListReq)

		productList, err := service.QueryProductList(ctx, req.StoreId)
		if err != nil {
			return nil, err
		}

		return productList, nil
	}
}
