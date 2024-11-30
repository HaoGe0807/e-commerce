package repo

import (
	"context"
	"e-commerce/service/domain/product/entity"
)

type CategoryRepo interface {
	CreateCategory(ctx context.Context, category *entity.CategoryEntity) error
	GetCategory(ctx context.Context, categoryId string) (*entity.CategoryEntity, error)
}
