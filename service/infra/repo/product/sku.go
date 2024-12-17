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
	SpuId      string  `gorm:"column:spu_id"`
	SkuId      string  `gorm:"column:sku_id;primary_key"`
	SkuName    string  `gorm:"column:sku_name"`
	SellAmount float64 `gorm:"column:sell_amount"`
	CostAmount float64 `gorm:"column:cost_amount"`
	Deleted    bool    `gorm:"column:deleted"`
	IsDefault  bool    `gorm:"column:is_default"`
	Code       string  `gorm:"column:code"`
	Currency   string  `gorm:"currency"`
}

func (model *skuModel) convertEntityToModel(skuEntity *entity.SkuEntity) {
	model.SpuId = skuEntity.SpuId
	model.SkuId = skuEntity.SkuId
	model.SkuName = skuEntity.SkuName
	model.CostAmount = skuEntity.CostAmount.Amount
	model.SellAmount = skuEntity.SellAmount.Amount
	model.Currency = skuEntity.SellAmount.Currency
	model.Deleted = skuEntity.Deleted
	model.IsDefault = skuEntity.IsDefault
	model.Code = skuEntity.Code
}

func (model *skuModel) convertModelToEntity() *entity.SkuEntity {
	return &entity.SkuEntity{
		SpuId:   model.SpuId,
		SkuId:   model.SkuId,
		SkuName: model.SkuName,
		SellAmount: ebus.Money{
			Amount:   model.SellAmount,
			Currency: model.Currency,
		},
		CostAmount: ebus.Money{
			Amount:   model.CostAmount,
			Currency: model.Currency,
		},
		Deleted:   model.Deleted,
		IsDefault: model.IsDefault,
		Code:      model.Code,
	}
}

// CreateSku create sku
func (s SkuInfoRepoImpl) CreateSku(ctx context.Context, sku *entity.SkuEntity) error {
	model := &skuModel{}
	model.convertEntityToModel(sku)
	return s.db.Table(s.tableName).Create(model).Error
}

// UpdateSku update sku
func (s SkuInfoRepoImpl) UpdateSku(ctx context.Context, sku *entity.SkuEntity) error {
	model := &skuModel{}
	model.convertEntityToModel(sku)
	return s.db.Table(s.tableName).Save(model).Error
}

// DeleteSku delete sku
func (s SkuInfoRepoImpl) DeleteSku(ctx context.Context, skuId string) error {
	return s.db.Table(s.tableName).Where("sku_id = ?", skuId).Update("deleted", true).Error
}

// DeleteSkuBySpuId delete sku by spu id
func (s SkuInfoRepoImpl) DeleteSkuBySpuId(ctx context.Context, spuId string) error {
	return s.db.Table(s.tableName).Where("spu_id = ?", spuId).Update("deleted", true).Error
}

// GetSkuList get sku list
func (s SkuInfoRepoImpl) GetSkuList(ctx context.Context) ([]*entity.SkuEntity, error) {
	modelList := make([]skuModel, 0)
	entityList := make([]*entity.SkuEntity, 0)

	s.db.Table(s.tableName).Where("deleted = ?", false).Find(&modelList)
	for _, model := range modelList {
		entityList = append(entityList, model.convertModelToEntity())
	}
	return entityList, nil
}

// GetSkuListBySpuId get sku list
func (s SkuInfoRepoImpl) GetSkuListBySpuId(ctx context.Context, spuId string) ([]*entity.SkuEntity, error) {
	modelList := make([]skuModel, 0)
	entityList := make([]*entity.SkuEntity, 0)

	s.db.Table(s.tableName).Where("spu_id = ?", spuId).Where("deleted = ?", false).Find(&modelList)

	for _, model := range modelList {
		entityList = append(entityList, model.convertModelToEntity())
	}
	return entityList, nil
}

func (s SkuInfoRepoImpl) SaveSkuListBySpuId(ctx context.Context, skus []*entity.SkuEntity, spuId string) error {
	skuBOList, err := s.GetSkuListBySpuId(ctx, spuId)
	if err != nil {
		return err
	}

	//对比skuBOList 和 skus ，若存在skuBOList存在skus中不存在的skuBO，则将这批skuBO的deleted字段置为true
	for _, skuBO := range skuBOList {
		if !isExist(skuBO, skus) {
			err = s.DeleteSku(ctx, skuBO.SkuId)
			if err != nil {
				return err
			}
		}
	}

	for _, skuEntity := range skus {
		err = s.UpdateSku(ctx, skuEntity)
		if err != nil {
			return err
		}
	}

	return nil
}

func isExist(skuBO *entity.SkuEntity, skus []*entity.SkuEntity) bool {
	for _, sku := range skus {
		if skuBO.SkuId == sku.SkuId {
			return true
		}
	}
	return false
}
