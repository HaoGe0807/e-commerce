package vo

import (
	"e-commerce/service/domain/product/entity"
	"e-commerce/service/infra/ebus"
)

type ProductVO struct {
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
	// sku info
	Skus []skuVO
}

type skuVO struct {
	StoreId      string
	SpuId        string
	SkuId        string
	SkuName      string
	SellAmount   ebus.Money
	CostAmount   ebus.Money
	Deleted      bool
	IsDefault    bool
	Code         string
	MinimumStock int64
	stock        int64
}

type ProductVOList []ProductVO

func ProductBOToVO(info entity.ProductAggInfo, inv []string) ProductVO {
	return ProductVO{
		SpuId:                 info.SpuId,
		CategoryId:            info.CategoryId,
		StoreId:               info.StoreId,
		ProductName:           info.ProductName,
		Unit:                  info.Unit,
		Status:                info.Status,
		MnemonicCode:          info.MnemonicCode,
		ProductSpecifications: info.ProductSpecifications,
		Icon:                  info.Icon,
		CustomizationList:     info.CustomizationList,
		IngredientList:        info.IngredientList,
		Deleted:               info.Deleted,
		PriceMethod:           info.PriceMethod,
		Sort:                  info.Sort,
		SortFiled:             info.SortFiled,
		Shape:                 info.Shape,
		ShapeColor:            info.ShapeColor,
		ProductType:           info.ProductType,
		Version:               info.Version,
		FirstDisplay:          info.FirstDisplay,
		Skus:                  skuBOToVO(info.Skus, inv),
	}
}

func skuBOToVO(skuEntityList []entity.SkuEntity, inv []string) []skuVO {
	skuListVO := make([]skuVO, 0)
	for _, skuEntity := range skuEntityList {
		skuListVO = append(skuListVO, skuVO{
			StoreId:      skuEntity.StoreId,
			SpuId:        skuEntity.SpuId,
			SkuId:        skuEntity.SkuId,
			SkuName:      skuEntity.SkuName,
			SellAmount:   skuEntity.SellAmount,
			CostAmount:   skuEntity.CostAmount,
			Deleted:      skuEntity.Deleted,
			IsDefault:    skuEntity.IsDefault,
			Code:         skuEntity.Code,
			MinimumStock: skuEntity.MinimumStock,
			stock:        0,
		})
	}
	return skuListVO
}
