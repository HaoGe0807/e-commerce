package entity

import (
	"context"
	"e-commerce/service/infra/ebus"
	"e-commerce/service/infra/utils"
)

type SkuEntity struct {
	SpuId      string     `json:"spu_id"`
	SkuId      string     `json:"sku_id"`
	SkuName    string     `json:"sku_name"`
	SellAmount ebus.Money `json:"sell_amount"`
	CostAmount ebus.Money `json:"cost_amount"`
	Deleted    bool       `json:"deleted"`
	IsDefault  bool       `json:"is_default"`
	Code       string     `json:"code"`
}

// 为skuEntity填充字段值
func (s *SkuEntity) FillField(ctx context.Context, spuId, skuName string, sellAmount, costAmount ebus.Money, isDefault bool, code string) *SkuEntity {
	if s.SkuId == "" {
		s.SkuId = utils.ModelIdNext("sku")
	}

	s.SpuId = spuId
	s.SkuName = skuName
	s.SellAmount = sellAmount
	s.CostAmount = costAmount
	s.Deleted = false
	s.IsDefault = isDefault
	s.Code = code
	return s
}
