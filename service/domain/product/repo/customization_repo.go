package repo

import (
	"context"
	"e-commerce/service/domain/product/entity"
)

type CustomizationRepo interface {
	CreateCustomization(ctx context.Context, customization *entity.CustomizationEntity) error
	GetCustomization(ctx context.Context, customizationId string) (*entity.CustomizationEntity, error)
}
