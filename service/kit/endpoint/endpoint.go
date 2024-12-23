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

	CreateCategoryEP    goEndpoint.Endpoint
	UpdateCategoryEP    goEndpoint.Endpoint
	DeleteCategoryEP    goEndpoint.Endpoint
	QueryCategoryEP     goEndpoint.Endpoint
	QueryCategoryListEP goEndpoint.Endpoint
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

	createCategoryEP := NewCreateCategoryEP(service)
	createCategoryEP = utils.LoggingMiddleware()(createCategoryEP)
	createCategoryEP = utils.ParameterCheckMiddleware()(createCategoryEP)

	updateCategoryEP := NewUpdateCategoryEP(service)
	updateCategoryEP = utils.LoggingMiddleware()(updateCategoryEP)
	updateCategoryEP = utils.ParameterCheckMiddleware()(updateCategoryEP)

	deleteCategoryEP := NewDeleteCategoryEP(service)
	deleteCategoryEP = utils.LoggingMiddleware()(deleteCategoryEP)
	deleteCategoryEP = utils.ParameterCheckMiddleware()(deleteCategoryEP)

	queryCategoryEP := NewQueryCategoryEP(service)
	queryCategoryEP = utils.LoggingMiddleware()(queryCategoryEP)
	queryCategoryEP = utils.ParameterCheckMiddleware()(queryCategoryEP)

	queryCategoryListEP := NewQueryCategoryListEP(service)
	queryCategoryListEP = utils.LoggingMiddleware()(queryCategoryListEP)
	queryCategoryListEP = utils.ParameterCheckMiddleware()(queryCategoryListEP)
	return Set{
		CreateProductEP:    createProductEP,
		UpdateProductEP:    updateProductEP,
		DeleteProductEP:    deleteProductEP,
		QueryProductEP:     queryProductEP,
		QueryProductListEP: queryProductListEP,
	}
}

// @Description 根据提供的信息创建一个新的产品
// @Tags 产品服务
// @Produce json
// @Param requestBody body CreateProductReq true "创建产品请求体，包含产品相关信息"
// @Success 201 {object} CreateProductResp "成功创建产品，返回创建后的产品信息"
// @Failure 400 {object} error "请求参数错误，例如请求体格式错误等"
// @Failure 404 {object} error "未找到相关资源或执行创建操作失败"
// @Router /api/e-commerce/product/createProduct [POST]
func NewCreateProductEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateProductReq)

		productAggInfo, err := service.CreateProduct(ctx, req.ProductName, req.CategoryId, req.Skus, req.Status, req.Icon)
		if err != nil {
			return nil, err
		}
		resp := productAggInfo
		return resp, nil
	}
}

// @Summary 更新产品
// @Description 根据提供的信息更新指定的产品
// @Tags 产品服务
// @Produce json
// @Param requestBody body UpdateProductReq true "更新产品请求体，包含更新后的产品相关信息"
// @Success 200 {object} UpdateProductResp "成功更新产品，返回更新后的产品信息"
// @Failure 400 {object} error "请求参数错误，例如请求体格式错误等"
// @Failure 404 {object} error "未找到要更新的产品或执行更新操作失败"
// @Router /api/e-commerce/product/updateProduct [POST]
func NewUpdateProductEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateProductReq)

		productAggInfo, err := service.UpdateProduct(ctx, req.SpuId, req.ProductName, req.CategoryId, req.Skus, req.Status, req.Icon)
		if err != nil {
			return nil, err
		}
		resp := productAggInfo
		return resp, nil
	}
}

// @Summary 删除产品
// @Description 根据提供的产品标识删除指定的产品
// @Tags 产品服务
// @Produce json
// @Param requestBody body DeleteProductReq true "删除产品请求体，包含产品标识等相关信息"
// @Success 200 {object} DeleteProductResp "成功删除产品，返回成功响应信息"
// @Failure 400 {object} error "请求参数错误，例如请求体格式错误等"
// @Failure 404 {object} error "未找到要删除的产品或执行删除操作失败"
// @Router /api/e-commerce/product/deleteProduct [POST]
func NewDeleteProductEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteProductReq)

		err = service.DeleteProduct(ctx, req.SpuId)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

// @Summary 查询单个产品
// @Description 根据提供的查询条件查询单个产品的信息
// @Tags 产品服务
// @Produce json
// @Param requestBody body QueryProductReq true "查询产品请求体，包含查询条件等相关信息"
// @Success 200 {object} vo.ProductVO "成功查询到产品，返回产品信息"
// @Failure 400 {object} error "请求参数错误，例如请求体格式错误等"
// @Failure 404 {object} error "未找到符合查询条件的产品"
// @Router /api/e-commerce/product/queryProduct [POST]
func NewQueryProductEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(QueryProductReq)

		product, err := service.QueryProduct(ctx, req.SpuId)
		if err != nil {
			return nil, err
		}

		return product, nil
	}
}

// @Summary 查询产品列表
// @Description 根据提供的查询条件查询产品的 List信息
// @Tags 产品服务
// @Produce json
// @Param noParams query string false "此接口无需传入参数，此参数仅为占位示意，实际不会使用。"
// @Success 200 {object} []vo.ProductVO "成功查询到产品 List，返回产品列表信息"
// @Failure 400 {object} error "请求参数错误，例如请求体格式错误等"
// @Failure 404 {object} error "未找到符合查询条件的产品列表"
// @Router /api/e-commerce/product/queryProductList [POST]
func NewQueryProductListEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		productList, err := service.QueryProductList(ctx)
		if err != nil {
			return nil, err
		}

		return productList, nil
	}
}

// category
func NewCreateCategoryEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateCategoryReq)

		categoryInfo, err := service.CreateCategory(ctx, req.CategoryName)
		if err != nil {
			return nil, err
		}

		return categoryInfo, nil
	}
}

func NewUpdateCategoryEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateCategoryReq)

		categoryInfo, err := service.UpdateCategory(ctx, req.CategoryId, req.CategoryName)
		if err != nil {
			return nil, err
		}

		return categoryInfo, nil
	}
}

func NewDeleteCategoryEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteCategoryReq)

		err = service.DeleteCategory(ctx, req.CategoryId)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func NewQueryCategoryEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(QueryCategoryReq)

		categoryInfo, err := service.QueryCategory(ctx, req.CategoryId)
		if err != nil {
			return nil, err
		}

		return categoryInfo, nil
	}
}

func NewQueryCategoryListEP(service app.ECommerceService) goEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		categoryList, err := service.QueryCategoryList(ctx)
		if err != nil {
			return nil, err
		}

		return categoryList, nil
	}
}
