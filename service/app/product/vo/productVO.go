package vo

import (
	"e-commerce/service/domain/product/entity"
	"e-commerce/service/infra/ebus"
)

type ProductVO struct {
	SpuId       string `json:"spu_id"`
	CategoryId  string `json:"category_id"`
	ProductName string `json:"product_name"`
	Status      string `json:"status"`
	Icon        string `json:"icon"`
	Deleted     bool   `json:"deleted"`
	// sku info
	Skus []*skuVO `json:"skus"`
}

type skuVO struct {
	SpuId      string     `json:"spu_id"`
	SkuId      string     `json:"sku_id"`
	SkuName    string     `json:"sku_name"`
	SellAmount ebus.Money `json:"sell_amount"`
	CostAmount ebus.Money `json:"cost_amount"`
	Deleted    bool       `json:"deleted"`
	IsDefault  bool       `json:"is_default"`
	Code       string     `json:"code"`
}

type ProductVOList []ProductVO

func ProductBOToVO(info *entity.ProductAggInfo) *ProductVO {
	return &ProductVO{
		SpuId:       info.SpuId,
		CategoryId:  info.CategoryId,
		ProductName: info.ProductName,
		Status:      info.Status,
		Icon:        info.Icon,
		Deleted:     info.Deleted,
		Skus:        skuBOToVO(info.Skus),
	}
}

func skuBOToVO(skuEntityList []*entity.SkuEntity) []*skuVO {
	skuListVO := make([]*skuVO, 0)
	for _, skuEntity := range skuEntityList {
		skuListVO = append(skuListVO, &skuVO{
			SpuId:      skuEntity.SpuId,
			SkuId:      skuEntity.SkuId,
			SkuName:    skuEntity.SkuName,
			SellAmount: skuEntity.SellAmount,
			CostAmount: skuEntity.CostAmount,
			Deleted:    skuEntity.Deleted,
			IsDefault:  skuEntity.IsDefault,
			Code:       skuEntity.Code,
		})
	}
	return skuListVO
}
