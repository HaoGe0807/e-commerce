package entity

import (
	"context"
	"e-commerce/service/infra/utils"
)

type SpuEntity struct {
	SpuId                 string
	CategoryId            string
	StoreId               string
	ProductName           string
	Unit                  string
	Status                string
	MnemonicCode          string
	ProductSpecifications string
	Icon                  string
	CustomizationList     string
	IngredientList        string
	Deleted               bool
	PriceMethod           string
	Sort                  int
	SortFiled             string
	Shape                 string
	ShapeColor            string
	ProductType           string
	Version               string
	FirstDisplay          string
}

// 为spuEntity填充字段值
func (s *SpuEntity) FillField(ctx context.Context, storeId string, productName string, unit string, categoryId string, mnemonicCode string, status string, icon string, priceMethod string, productSpecifications string, shape string, shapeColor string, firstDisplay string, productType string) *SpuEntity {
	if s.SpuId == "" {
		s.SpuId = utils.ModelIdNext("spu")
	}

	s.StoreId = storeId
	s.ProductName = productName
	s.Unit = unit
	s.CategoryId = categoryId
	s.MnemonicCode = mnemonicCode
	s.Status = status
	s.Icon = icon
	s.PriceMethod = priceMethod
	s.ProductSpecifications = productSpecifications
	s.Shape = shape
	s.ShapeColor = shapeColor
	s.FirstDisplay = firstDisplay
	s.ProductType = productType
	return s
}
