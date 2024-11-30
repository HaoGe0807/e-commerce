package product

import (
	"context"
	"e-commerce/service/domain/product/entity"
	"e-commerce/service/domain/product/repo"
	"e-commerce/service/infra/consts"
	"e-commerce/service/infra/orm"
	"github.com/jinzhu/gorm"
)

var CategoryTableName = "product_category"

var CategoryRepoImpl = &CategoryInfoRepoImpl{}

type CategoryInfoRepoImpl struct {
	db        *gorm.DB
	tableName string
}

func NewCategoryRepo() repo.CategoryRepo {
	return &CategoryInfoRepoImpl{
		db:        orm.GetORM(consts.DB_RETAIL),
		tableName: CategoryTableName,
	}
}

type categoryModel struct {
	StoreId      string `gorm:"column:store_id"`
	CategoryId   string `gorm:"column:category_id"`
	CategoryName string `gorm:"column:category_name"`
	Deleted      bool   `gorm:"column:deleted"`
}

func (model categoryModel) convertEntityToModel(skuEntity *entity.CategoryEntity) {
	model.StoreId = skuEntity.StoreId
	model.CategoryId = skuEntity.CategoryId
	model.CategoryName = skuEntity.CategoryName
	model.Deleted = skuEntity.Deleted
}

// model转entity
func (model categoryModel) convertModelToEntity() *entity.CategoryEntity {
	return &entity.CategoryEntity{
		StoreId:      model.StoreId,
		CategoryId:   model.CategoryId,
		CategoryName: model.CategoryName,
		Deleted:      model.Deleted,
	}
}

func (s CategoryInfoRepoImpl) CreateCategory(ctx context.Context, category *entity.CategoryEntity) error {
	model := categoryModel{}
	model.convertEntityToModel(category)
	return s.db.Table(s.tableName).Create(&model).Error
}

func (s CategoryInfoRepoImpl) GetCategory(ctx context.Context, categoryId string) (*entity.CategoryEntity, error) {
	model := categoryModel{}
	s.db.Table(s.tableName).Where("category_id = ?", categoryId).First(&model)

	return model.convertModelToEntity(), nil
}
