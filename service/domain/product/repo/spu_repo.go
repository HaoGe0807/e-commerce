package repo

import (
	"context"
	"ge-commerce/service/domain/product/entity"
)

type SpuRepo interface {
	CreateSpu(ctx context.Context, spu *entity.SpuEntity) error
	UpdateSpu(ctx context.Context, spu *entity.SpuEntity) error
	DeleteSpu(ctx context.Context, storeId, spuId string) error
	GetSpu(ctx context.Context, storeId, spuId string) (*entity.SpuEntity, error)
	GetSpuListByStoreId(ctx context.Context, storeId string) ([]*entity.SpuEntity, error)
}
