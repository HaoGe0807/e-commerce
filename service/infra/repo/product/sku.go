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
		db:        orm.GetORM(consts.DB_NAME),
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
	model.SkuId = skuEntity.SkuId
	model.SkuName = skuEntity.SkuName
	model.SellAmount = skuEntity.SellAmount
	model.CostAmount = skuEntity.CostAmount
	model.Deleted = skuEntity.Deleted
	model.IsDefault = skuEntity.IsDefault
	model.Code = skuEntity.Code
}

func (model skuModel) convertModelToEntity() *entity.SkuEntity {
	return &entity.SkuEntity{
		SpuId:      model.SpuId,
		SkuId:      model.SkuId,
		SkuName:    model.SkuName,
		SellAmount: model.SellAmount,
		CostAmount: model.CostAmount,
		Deleted:    model.Deleted,
		IsDefault:  model.IsDefault,
		Code:       model.Code,
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
func (s SkuInfoRepoImpl) DeleteSku(ctx context.Context, skuId string) error {
	return s.db.Table(s.tableName).Where("sku_id = ?", skuId).Update("deleted", false).Error
}

// DeleteSkuBySpuId delete sku by spu id
func (s SkuInfoRepoImpl) DeleteSkuBySpuId(ctx context.Context, spuId string) error {
	return s.db.Table(s.tableName).Where("spu_id = ?", spuId).Update("deleted", false).Error
}

// GetSkuList get sku list
func (s SkuInfoRepoImpl) GetSkuList(ctx context.Context) ([]*entity.SkuEntity, error) {
	modelList := make([]skuModel, 0)
	entityList := make([]*entity.SkuEntity, 0)

	s.db.Table(s.tableName).Where("deleted", false).Find(&modelList)
	for _, model := range modelList {
		entityList = append(entityList, model.convertModelToEntity())
	}
	return entityList, nil
}

// GetSkuListBySpuId get sku list
func (s SkuInfoRepoImpl) GetSkuListBySpuId(ctx context.Context, spuId string) ([]*entity.SkuEntity, error) {
	modelList := make([]skuModel, 0)
	entityList := make([]*entity.SkuEntity, 0)

	s.db.Table(s.tableName).Where("spu_id", spuId).Find(&modelList)
	for _, model := range modelList {
		entityList = append(entityList, model.convertModelToEntity())
	}
	return entityList, nil
}
