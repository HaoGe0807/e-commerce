package product

import (
	"context"
	"e-commerce/service/domain/product/entity"
	"e-commerce/service/domain/product/repo"
	"e-commerce/service/infra/consts"
	"e-commerce/service/infra/ebus"
	"e-commerce/service/infra/orm"
	"github.com/jinzhu/gorm"
)

var SkuTableName = "sku"

var SkuRepoImpl = &SkuInfoRepoImpl{}

type SkuInfoRepoImpl struct {
	db        *gorm.DB
	tableName string
}

func NewSkuRepo() repo.SkuRepo {
	return &SkuInfoRepoImpl{
		db:        orm.GetORM(consts.DB_RETAIL),
		tableName: SkuTableName,
	}
}

type skuModel struct {
	SpuId        string     `gorm:"column:spu_id"`
	StoreId      string     `gorm:"column:store_id"`
	SkuId        string     `gorm:"column:sku_id"`
	SkuName      string     `gorm:"column:sku_name"`
	SellAmount   ebus.Money `gorm:"column:sell_amount"`
	CostAmount   ebus.Money `gorm:"column:cost_amount"`
	Deleted      bool       `gorm:"column:deleted"`
	IsDefault    bool       `gorm:"column:is_default"`
	Code         string     `gorm:"column:code"`
	MinimumStock int64      `gorm:"column:minimum_stock"`
}

func (model skuModel) convertEntityToModel(skuEntity *entity.SkuEntity) {
	model.SpuId = skuEntity.SpuId
	model.StoreId = skuEntity.StoreId
	model.SkuId = skuEntity.SkuId
	model.SkuName = skuEntity.SkuName
	model.SellAmount = skuEntity.SellAmount
	model.CostAmount = skuEntity.CostAmount
	model.Deleted = skuEntity.Deleted
	model.IsDefault = skuEntity.IsDefault
	model.Code = skuEntity.Code
	model.MinimumStock = skuEntity.MinimumStock
}

func (model skuModel) convertModelToEntity() *entity.SkuEntity {
	return &entity.SkuEntity{
		StoreId:      model.StoreId,
		SpuId:        model.SpuId,
		SkuId:        model.SkuId,
		SkuName:      model.SkuName,
		SellAmount:   model.SellAmount,
		CostAmount:   model.CostAmount,
		Deleted:      model.Deleted,
		IsDefault:    model.IsDefault,
		Code:         model.Code,
		MinimumStock: model.MinimumStock,
	}
}

// CreateSku create sku
func (s SkuInfoRepoImpl) CreateSku(ctx context.Context, sku *entity.SkuEntity) error {
	model := skuModel{}
	model.convertEntityToModel(sku)
	return s.db.Table(s.tableName).Create(&model).Error
}

// UpdateSku update sku
func (s SkuInfoRepoImpl) UpdateSku(ctx context.Context, sku *entity.SkuEntity) error {
	model := skuModel{}
	model.convertEntityToModel(sku)
	return s.db.Table(s.tableName).Save(&model).Error
}

// DeleteSku delete sku
func (s SkuInfoRepoImpl) DeleteSku(ctx context.Context, storeId, skuId string) error {
	return s.db.Table(s.tableName).Where("sku_id = ?", skuId).Update("deleted", true).Error
}

// DeleteSkuBySpuId delete sku by spu id
func (s SkuInfoRepoImpl) DeleteSkuBySpuId(ctx context.Context, storeId, spuId string) error {
	return s.db.Table(s.tableName).Where("spu_id = ?", spuId).Update("deleted", true).Error
}

// GetSkuListByStoreId get sku list
func (s SkuInfoRepoImpl) GetSkuListByStoreId(ctx context.Context, storeId string) ([]*entity.SkuEntity, error) {
	modelList := make([]skuModel, 0)
	entityList := make([]*entity.SkuEntity, 0)

	s.db.Table(s.tableName).Where("store_id = ?", storeId).Find(&modelList)
	for _, model := range modelList {
		entityList = append(entityList, model.convertModelToEntity())
	}
	return entityList, nil
}

// GetSkuListByStoreIdAndSpuId get sku list
func (s SkuInfoRepoImpl) GetSkuListByStoreIdAndSpuId(ctx context.Context, storeId, spuId string) ([]*entity.SkuEntity, error) {
	modelList := make([]skuModel, 0)
	entityList := make([]*entity.SkuEntity, 0)

	s.db.Table(s.tableName).Where("store_id = ?", storeId).Where("spu_id", spuId).Find(&modelList)
	for _, model := range modelList {
		entityList = append(entityList, model.convertModelToEntity())
	}
	return entityList, nil
}