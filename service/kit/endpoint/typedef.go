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
	SkuName      string     `json:"sku_name" validate:"required"`
	SellAmount   ebus.Money `json:"sell_amount" validate:"required"`
	CostAmount   ebus.Money `json:"cost_amount" validate:"required"`
	IsDefault    bool       `json:"is_default" validate:"required"`
	Code         string     `json:"code" validate:"required"`
	Stock        int64      `json:"stock" validate:"required"`
	MinimumStock int64      `json:"minimum_stock" validate:"required"`
}

type CreateProductReq struct {
	ProductName           string                              `json:"product_name" validate:"required"`
	CategoryId            string                              `json:"category_id" validate:"required"`
	Skus                  []productEntity.SkuEntity           `json:"skus" validate:"required"`
	Unit                  string                              `json:"unit" validate:"required"`
	MnemonicCode          string                              `json:"mnemonic_code" validate:"required"`
	Status                string                              `json:"status" validate:"required"`
	StoreId               string                              `json:"store_id" validate:"required"`
	CustomizationList     []productEntity.CustomizationEntity `json:"customization_list"`
	IngredientList        []productEntity.IngredientEntity    `json:"ingredient_list"`
	Icon                  string                              `json:"icon"`
	PriceMethod           string                              `json:"price_method"`
	Shape                 string                              `json:"shape"`
	ShapeColor            string                              `json:"shape_color"`
	FirstDisplay          string                              `json:"first_display"`
	ProductSpecifications string                              `json:"product_specifications"`
	ProductType           string                              `json:"product_type"`
}

type CreateProductResp struct {
}

type UpdateProductReq struct {
	SpuId                 string                              `json:"spu_id" validate:"required"`
	ProductName           string                              `json:"product_name" validate:"required"`
	CategoryId            string                              `json:"category_id" validate:"required"`
	Skus                  []productEntity.SkuEntity           `json:"skus" validate:"required"`
	Unit                  string                              `json:"unit" validate:"required"`
	MnemonicCode          string                              `json:"mnemonic_code" validate:"required"`
	Status                string                              `json:"status" validate:"required"`
	StoreId               string                              `json:"store_id" validate:"required"`
	CustomizationList     []productEntity.CustomizationEntity `json:"customization_list"`
	IngredientList        []productEntity.IngredientEntity    `json:"ingredient_list"`
	Icon                  string                              `json:"icon"`
	PriceMethod           string                              `json:"price_method"`
	Shape                 string                              `json:"shape"`
	ShapeColor            string                              `json:"shape_color"`
	FirstDisplay          string                              `json:"first_display"`
	ProductSpecifications string                              `json:"product_specifications"`
	ProductType           string                              `json:"product_type"`
}

type UpdateProductResp struct{}

type DeleteProductReq struct {
	StoreId string `json:"store_id" validate:"required"`
	SpuId   string `json:"spu_id" validate:"required"`
}

type DeleteProductResp struct{}

type QueryProductReq struct {
	StoreId string `json:"store_id" validate:"required"`
	SpuId   string `json:"spu_id" validate:"required"`
}

type QueryProductResp productVO.ProductVO

type QueryProductListReq struct {
	StoreId string `json:"store_id" validate:"required"`
}

type QueryProductListResp []productVO.ProductVO
