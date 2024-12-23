package repo

import (
	"context"
	"e-commerce/service/domain/product/entity"
)

type CategoryRepo interface {
	CreateCategory(ctx context.Context, category *entity.CategoryEntity) error
	GetCategory(ctx context.Context, categoryId string) (*entity.CategoryEntity, error)
	GetCategoryByName(ctx context.Context, categoryId, categoryName string) (*entity.CategoryEntity, error)
	UpdateCategory(ctx context.Context, category *entity.CategoryEntity) error
	DeleteCategory(ctx context.Context, categoryId string) error
	GetCategoryList(ctx context.Context) ([]*entity.CategoryEntity, error)
}
