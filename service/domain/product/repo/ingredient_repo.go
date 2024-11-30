package repo

import (
	"context"
	"e-commerce/service/domain/product/entity"
)

type IngredientRepo interface {
	CreateIngredient(ctx context.Context, ingredient *entity.IngredientEntity) error
	GetIngredient(ctx context.Context, ingredient string) (*entity.IngredientEntity, error)
}
