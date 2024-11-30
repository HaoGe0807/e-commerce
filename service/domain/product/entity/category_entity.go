package entity

import (
	"context"
	"e-commerce/service/infra/utils"
)

type CategoryEntity struct {
	StoreId      string
	CategoryId   string
	CategoryName string
	Deleted      bool
}

// 为CategoryEntity填充字段值
func (s *CategoryEntity) FillField(ctx context.Context, storeId string, categoryName string, deleted bool) *CategoryEntity {
	if s.CategoryId == "" {
		s.CategoryId = utils.ModelIdNext("CategoryId")
	}

	s.StoreId = storeId
	s.CategoryName = categoryName
	s.Deleted = deleted
	return s
}
