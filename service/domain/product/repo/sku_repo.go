package repo

import (
	"context"
	"e-commerce/service/domain/product/entity"
)

type SkuRepo interface {
	CreateSku(ctx context.Context, sku *entity.SkuEntity) error
	UpdateSku(ctx context.Context, sku *entity.SkuEntity) error
	DeleteSku(ctx context.Context, storeId, skuId string) error
	DeleteSkuBySpuId(ctx context.Context, storeId, spuId string) error
	GetSkuListByStoreId(ctx context.Context, storeId string) ([]*entity.SkuEntity, error)
	GetSkuListByStoreIdAndSpuId(ctx context.Context, storeId, spuId string) ([]*entity.SkuEntity, error)
}
