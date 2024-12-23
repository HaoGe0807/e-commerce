package repo

import (
	"context"
	"e-commerce/service/domain/product/entity"
)

type SpuRepo interface {
	CreateSpu(ctx context.Context, spu *entity.SpuEntity) error
	UpdateSpu(ctx context.Context, spu *entity.SpuEntity) error
	DeleteSpu(ctx context.Context, spuId string) error
	GetSpu(ctx context.Context, spuId string) (*entity.SpuEntity, error)
	GetSpuList(ctx context.Context) ([]*entity.SpuEntity, error)
	GetSpuListByCategoryId(ctx context.Context, categoryId string) ([]*entity.SpuEntity, error)
}
