package entity

import (
	"context"
	"e-commerce/service/infra/utils"
)

type SpuEntity struct {
	SpuId       string `json:"spu_id"`
	CategoryId  string `json:"category_id"`
	ProductName string `json:"product_name"`
	Status      string `json:"status"`
	Icon        string `json:"icon"`
	Deleted     bool   `json:"deleted"`
}

// 为spuEntity填充字段值
func (s *SpuEntity) FillField(ctx context.Context, productName string, categoryId string, status string, icon string) *SpuEntity {
	if s.SpuId == "" {
		s.SpuId = utils.ModelIdNext("spu")
	}

	s.ProductName = productName
	s.CategoryId = categoryId
	s.Status = status
	s.Icon = icon
	return s
}
