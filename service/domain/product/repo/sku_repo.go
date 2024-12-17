package repo

import (
	"context"
	"e-commerce/service/domain/product/entity"
)

type SkuRepo interface {
	CreateSku(ctx context.Context, sku *entity.SkuEntity) error
	UpdateSku(ctx context.Context, sku *entity.SkuEntity) error
	SaveSkuListBySpuId(ctx context.Context, skus []*entity.SkuEntity, spuId string) error
	DeleteSku(ctx context.Context, skuId string) error
	DeleteSkuBySpuId(ctx context.Context, spuId string) error
	GetSkuListBySpuId(ctx context.Context, spuId string) ([]*entity.SkuEntity, error)
	GetSkuList(ctx context.Context) ([]*entity.SkuEntity, error)
}
