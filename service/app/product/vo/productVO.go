package vo

import (
	"e-commerce/service/domain/product/entity"
	"e-commerce/service/infra/ebus"
)

type ProductVO struct {
	SpuId       string
	CategoryId  string
	ProductName string
	Status      string
	Icon        string
	Deleted     bool
	// sku info
	Skus []skuVO
}

type skuVO struct {
	SpuId      string
	SkuId      string
	SkuName    string
	SellAmount ebus.Money
	CostAmount ebus.Money
	Deleted    bool
	IsDefault  bool
	Code       string
	stock      int64
}

type ProductVOList []ProductVO

func ProductBOToVO(info entity.ProductAggInfo, inv []string) ProductVO {
	return ProductVO{
		SpuId:       info.SpuId,
		CategoryId:  info.CategoryId,
		ProductName: info.ProductName,
		Status:      info.Status,
		Icon:        info.Icon,
		Deleted:     info.Deleted,
		Skus:        skuBOToVO(info.Skus, inv),
	}
}

func skuBOToVO(skuEntityList []entity.SkuEntity, inv []string) []skuVO {
	skuListVO := make([]skuVO, 0)
	for _, skuEntity := range skuEntityList {
		skuListVO = append(skuListVO, skuVO{
			SpuId:      skuEntity.SpuId,
			SkuId:      skuEntity.SkuId,
			SkuName:    skuEntity.SkuName,
			SellAmount: skuEntity.SellAmount,
			CostAmount: skuEntity.CostAmount,
			Deleted:    skuEntity.Deleted,
			IsDefault:  skuEntity.IsDefault,
			Code:       skuEntity.Code,
			stock:      0,
		})
	}
	return skuListVO
}
