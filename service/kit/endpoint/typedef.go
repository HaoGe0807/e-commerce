package endpoint

import (
	"e-commerce/service/domain/product/entity"
)

type Response interface {
	ToJson() ([]byte, error)
}

type CreateProductReq struct {
	ProductName string             `json:"product_name" validate:"required"`
	CategoryId  string             `json:"category_id" validate:"required"`
	Skus        []entity.SkuEntity `json:"skus" validate:"required"`
	Icon        string             `json:"icon"`
	Status      string             `json:"status" validate:"required"`
}

type UpdateProductReq struct {
	SpuId       string             `json:"spu_id" validate:"required"`
	ProductName string             `json:"product_name" validate:"required"`
	CategoryId  string             `json:"category_id" validate:"required"`
	Skus        []entity.SkuEntity `json:"skus" validate:"required"`
	Status      string             `json:"status" validate:"required"`
	Icon        string             `json:"icon"`
}

type DeleteProductReq struct {
	SpuId string `json:"spu_id" validate:"required"`
}

type QueryProductReq struct {
	SpuId string `json:"spu_id" validate:"required"`
}

type QueryProductListReq struct {
}

type CreateCategoryReq struct {
	CategoryName string `json:"category_name" validate:"required"`
}

type UpdateCategoryReq struct {
	CategoryId   string `json:"category_id" validate:"required"`
	CategoryName string `json:"category_name" validate:"required"`
}

type DeleteCategoryReq struct {
	CategoryId string `json:"category_id" validate:"required"`
}

type QueryCategoryReq struct {
	CategoryId string `json:"category_id" validate:"required"`
}

type QueryCategoryListReq struct{}
