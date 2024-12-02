package product

import (
	"context"
	"e-commerce/service/domain/product/entity"
	"e-commerce/service/domain/product/repo"
	"e-commerce/service/infra/consts"
	"e-commerce/service/infra/orm"
	"github.com/jinzhu/gorm"
)

var SpuTableName = "spu"

var SpuRepoImpl = &SpuInfoRepoImpl{}

type SpuInfoRepoImpl struct {
	db        *gorm.DB
	tableName string
}

func NewSpuRepo() repo.SpuRepo {
	return &SpuInfoRepoImpl{
		db:        orm.GetORM(consts.DB_NAME),
		tableName: SpuTableName,
	}
}

type spuModel struct {
	SpuId                 string `gorm:"column:spu_id"`
	CategoryId            string `gorm:"column:category_id"`
	StoreId               string `gorm:"column:store_id"`
	ProductName           string `gorm:"column:product_name"`
	Unit                  string `gorm:"column:unit"`
	Status                string `gorm:"column:status"`
	MnemonicCode          string `gorm:"column:mnemonic_code"`
	ProductSpecifications string `gorm:"column:product_specifications"`
	Icon                  string `gorm:"column:icon"`
	Deleted               bool   `gorm:"column:deleted"`
	PriceMethod           string `gorm:"column:price_method"`
	Sort                  int    `gorm:"column:sort"`
	SortFiled             string `gorm:"column:sort_filed"`
	Shape                 string `gorm:"column:shape"`
	ShapeColor            string `gorm:"column:shape_color"`
	FistDisplay           string `gorm:"column:first_display"`
	ProductType           string `gorm:"column:product_type"`
	Version               string `gorm:"column:version"`
}

func (model spuModel) convertEntityToModel(spuEntity *entity.SpuEntity) {
	model.SpuId = spuEntity.SpuId
	model.CategoryId = spuEntity.CategoryId
	model.StoreId = spuEntity.StoreId
	model.ProductName = spuEntity.ProductName
	model.Unit = spuEntity.Unit
	model.Status = spuEntity.Status
	model.MnemonicCode = spuEntity.MnemonicCode
	model.ProductSpecifications = spuEntity.ProductSpecifications
	model.Icon = spuEntity.Icon
	model.Deleted = spuEntity.Deleted
	model.PriceMethod = spuEntity.PriceMethod
	model.Sort = spuEntity.Sort
	model.SortFiled = spuEntity.SortFiled
	model.Shape = spuEntity.Shape
}

func (model spuModel) convertModelToEntity() *entity.SpuEntity {
	return &entity.SpuEntity{
		SpuId:                 model.SpuId,
		CategoryId:            model.CategoryId,
		StoreId:               model.StoreId,
		ProductName:           model.ProductName,
		Unit:                  model.Unit,
		Status:                model.Status,
		MnemonicCode:          model.MnemonicCode,
		ProductSpecifications: model.ProductSpecifications,
		Icon:                  model.Icon,
		Deleted:               model.Deleted,
		PriceMethod:           model.PriceMethod,
		Sort:                  model.Sort,
		SortFiled:             model.SortFiled,
		Shape:                 model.Shape,
	}
}

// CreateSpu 创建
func (s SpuInfoRepoImpl) CreateSpu(ctx context.Context, spu *entity.SpuEntity) error {
	model := spuModel{}
	model.convertEntityToModel(spu)
	return s.db.Table(s.tableName).Create(&model).Error
}

// UpdateSpu 更新
func (s SpuInfoRepoImpl) UpdateSpu(ctx context.Context, spu *entity.SpuEntity) error {
	model := spuModel{}
	model.convertEntityToModel(spu)
	return s.db.Table(s.tableName).Save(&model).Error
}

// DeleteSpu 删除
func (s SpuInfoRepoImpl) DeleteSpu(ctx context.Context, storeId, spuId string) error {
	return s.db.Table(s.tableName).Where("spu_id = ?", spuId).Update("deleted", true).Error
}

// GetSpu 获取
func (s SpuInfoRepoImpl) GetSpu(ctx context.Context, storeId, spuId string) (*entity.SpuEntity, error) {
	model := spuModel{}
	s.db.Table(s.tableName).Where("spu_id = ?", spuId).First(&model)
	return model.convertModelToEntity(), nil
}

// GetSpuList 获取列表
func (s SpuInfoRepoImpl) GetSpuListByStoreId(ctx context.Context, storeId string) ([]*entity.SpuEntity, error) {
	modelList := make([]spuModel, 0)
	entityList := make([]*entity.SpuEntity, 0)

	s.db.Table(s.tableName).Where("store_id = ?", storeId).Find(&modelList)
	for _, model := range modelList {
		entityList = append(entityList, model.convertModelToEntity())
	}
	return entityList, nil
}

// GetSpuListByCategoryId 获取列表
func (s SpuInfoRepoImpl) GetSpuListByCategoryId(ctx context.Context, storeId, categoryId string) ([]*entity.SpuEntity, error) {
	modelList := make([]spuModel, 0)
	entityList := make([]*entity.SpuEntity, 0)

	s.db.Table(s.tableName).Where("category_id = ?", categoryId).Find(&modelList)
	for _, model := range modelList {
		entityList = append(entityList, model.convertModelToEntity())
	}
	return entityList, nil
}
