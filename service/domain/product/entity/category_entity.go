package entity

import (
	"context"
	"e-commerce/service/infra/utils"
)

type CategoryEntity struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	Deleted      bool   `json:"deleted"`
}

// 为CategoryEntity填充字段值
func (s *CategoryEntity) FillField(ctx context.Context, categoryName string, deleted bool) *CategoryEntity {
	if s.CategoryId == "" {
		s.CategoryId = utils.ModelIdNext("CategoryId")
	}

	s.CategoryName = categoryName
	s.Deleted = deleted
	return s
}
