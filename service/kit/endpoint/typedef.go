package endpoint

import (
	productVO "e-commerce/service/app/product/vo"
	"e-commerce/service/domain/product/entity"
	"encoding/json"
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

type CreateProductResp struct {
	*productVO.ProductVO
}

func (resp *CreateProductResp) ToJson() ([]byte, error) {
	return json.Marshal(resp)
}

type UpdateProductReq struct {
	SpuId       string             `json:"spu_id" validate:"required"`
	ProductName string             `json:"product_name" validate:"required"`
	CategoryId  string             `json:"category_id" validate:"required"`
	Skus        []entity.SkuEntity `json:"skus" validate:"required"`
	Status      string             `json:"status" validate:"required"`
	Icon        string             `json:"icon"`
}

type UpdateProductResp struct {
	*productVO.ProductVO
}

func (resp *UpdateProductResp) ToJson() ([]byte, error) {
	return json.Marshal(resp)
}

type DeleteProductReq struct {
	SpuId string `json:"spu_id" validate:"required"`
}

type DeleteProductResp struct{}

type QueryProductReq struct {
	SpuId string `json:"spu_id" validate:"required"`
}

type QueryProductResp struct {
	*productVO.ProductVO
}

func (resp *QueryProductResp) ToJson() ([]byte, error) {
	return json.Marshal(resp)
}

type QueryProductListReq struct {
}

type QueryProductListResp struct {
	Products *[]productVO.ProductVO
}

func (resp *QueryProductListResp) ToJson() ([]byte, error) {
	return json.Marshal(resp)
}
