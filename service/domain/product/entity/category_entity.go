package entity

import (
	"context"
	"e-commerce/service/infra/log"
	"e-commerce/service/infra/utils"
	"encoding/json"
)

type CategoryEntity struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	Deleted      bool   `json:"deleted"`
}

// 为CategoryEntity填充字段值
func (s *CategoryEntity) FillField(ctx context.Context, categoryName string) *CategoryEntity {
	if s.CategoryId == "" {
		s.CategoryId = utils.ModelIdNext("CategoryId")
	}

	s.CategoryName = categoryName
	s.Deleted = false
	return s
}

func ConvertStringToCategoryInfo(str string) (*CategoryEntity, error) {
	categoryInfo := &CategoryEntity{}
	err := json.Unmarshal([]byte(str), categoryInfo)
	if err != nil {
		log.Error("Error converting Redis string to Spu struct: %v\n", err)
		log.Error("The Redis string content is: %s\n", str)
		return &CategoryEntity{}, err
	}

	return categoryInfo, nil
}

func CategoryInfoToJsonMarshal(categoryEntity *CategoryEntity) []byte {
	marshal, err := json.Marshal(categoryEntity)
	if err != nil {
		return nil
	}
	return marshal
}
