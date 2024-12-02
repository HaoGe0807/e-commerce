package entity

import (
	"context"
	"e-commerce/service/infra/utils"
)

type SpuEntity struct {
	SpuId       string
	CategoryId  string
	ProductName string
	Status      string
	Icon        string
	Deleted     bool
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
