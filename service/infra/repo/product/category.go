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
		db:        orm.GetORM(consts.DB_NAME),
		tableName: CategoryTableName,
	}
}

type categoryModel struct {
	CategoryId   string `gorm:"column:category_id;primary_key"`
	CategoryName string `gorm:"column:category_name"`
	Deleted      bool   `gorm:"column:deleted"`
}

func (model categoryModel) convertEntityToModel(skuEntity *entity.CategoryEntity) {
	model.CategoryId = skuEntity.CategoryId
	model.CategoryName = skuEntity.CategoryName
	model.Deleted = skuEntity.Deleted
}

// model转entity
func (model categoryModel) convertModelToEntity() *entity.CategoryEntity {
	return &entity.CategoryEntity{
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

func (s CategoryInfoRepoImpl) UpdateCategory(ctx context.Context, category *entity.CategoryEntity) error {
	model := categoryModel{}
	model.convertEntityToModel(category)
	return s.db.Table(s.tableName).Save(&model).Error
}

func (s CategoryInfoRepoImpl) GetCategory(ctx context.Context, categoryId string) (*entity.CategoryEntity, error) {
	model := categoryModel{}
	s.db.Table(s.tableName).Where("category_id = ?", categoryId).First(&model)

	return model.convertModelToEntity(), nil
}

func (s CategoryInfoRepoImpl) DeleteCategory(ctx context.Context, categoryId string) error {
	return s.db.Table(s.tableName).Where("category_id = ?", categoryId).Update("deleted", true).Error
}

func (s CategoryInfoRepoImpl) GetCategoryList(ctx context.Context) ([]*entity.CategoryEntity, error) {
	modelList := make([]categoryModel, 0)
	entityList := make([]*entity.CategoryEntity, 0)

	s.db.Table(s.tableName).Where("deleted = ?", false).Find(&modelList)
	for _, model := range modelList {
		entityList = append(entityList, model.convertModelToEntity())
	}
	return entityList, nil
}

func (s CategoryInfoRepoImpl) GetCategoryByName(ctx context.Context, categoryId, categoryName string) (*entity.CategoryEntity, error) {
	model := categoryModel{}
	sql := s.db.Table(s.tableName).Where("category_name = ?", categoryName).Where("deleted = ?", false)

	if categoryId != "" {
		sql.Where("category_id != ?", categoryId)

	}

	sql.First(&model)

	return model.convertModelToEntity(), nil
}
