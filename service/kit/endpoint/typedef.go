package endpoint

import (
	productVO "e-commerce/service/app/product/vo"
	productEntity "e-commerce/service/domain/product/entity"
	"e-commerce/service/infra/ebus"
)

type Response interface {
	ToJson() ([]byte, error)
}

type Sku struct {
	SkuName    string     `json:"sku_name" validate:"required"`
	SellAmount ebus.Money `json:"sell_amount" validate:"required"`
	CostAmount ebus.Money `json:"cost_amount" validate:"required"`
	IsDefault  bool       `json:"is_default" validate:"required"`
	Code       string     `json:"code" validate:"required"`
	Stock      int64      `json:"stock" validate:"required"`
}

type CreateProductReq struct {
	ProductName string                    `json:"product_name" validate:"required"`
	CategoryId  string                    `json:"category_id" validate:"required"`
	Skus        []productEntity.SkuEntity `json:"skus" validate:"required"`
	Icon        string                    `json:"icon"`
	Status      string                    `json:"status" validate:"required"`
}

type CreateProductResp struct {
}

type UpdateProductReq struct {
	SpuId       string                    `json:"spu_id" validate:"required"`
	ProductName string                    `json:"product_name" validate:"required"`
	CategoryId  string                    `json:"category_id" validate:"required"`
	Skus        []productEntity.SkuEntity `json:"skus" validate:"required"`
	Status      string                    `json:"status" validate:"required"`
	Icon        string                    `json:"icon"`
}

type UpdateProductResp struct{}

type DeleteProductReq struct {
	SpuId string `json:"spu_id" validate:"required"`
}

type DeleteProductResp struct{}

type QueryProductReq struct {
	SpuId string `json:"spu_id" validate:"required"`
}

type QueryProductResp productVO.ProductVO

type QueryProductListReq struct {
}

type QueryProductListResp []productVO.ProductVO
