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

var IngredientTableName = "ingredient"

var IngredientTableNameRepoImpl = &IngredientInfoRepoImpl{}

type IngredientInfoRepoImpl struct {
	db        *gorm.DB
	tableName string
}

func NewIngredientRepo() repo.IngredientRepo {
	return &IngredientInfoRepoImpl{
		db:        orm.GetORM(consts.DB_NAME),
		tableName: CustomizationTableName,
	}
}

type ingredientModel struct {
	StoreId           string     `gorm:"column:store_id"`
	IngredientId      string     `gorm:"column:ingredient_id"`
	IngredientName    string     `gorm:"column:ingredient_name"`
	Deleted           bool       `gorm:"column:deleted"`
	IngredientGroupId string     `gorm:"column:ingredient_group_id"`
	Price             ebus.Money `gorm:"column:price"`
	MarkUp            bool       `gorm:"column:mark_up"`
	deleted           bool       `gorm:"column:deleted"`
}

func (model ingredientModel) convertEntityToModel(entity *entity.IngredientEntity) {
	model.StoreId = entity.StoreId
	model.IngredientId = entity.IngredientId
	model.IngredientName = entity.IngredientName
	model.IngredientGroupId = entity.IngredientGroupId
	model.Price = entity.Price
	model.MarkUp = entity.MarkUp
	model.deleted = entity.Deleted
}

// model转entity
func (model ingredientModel) convertModelToEntity() *entity.IngredientEntity {
	return &entity.IngredientEntity{
		StoreId:           model.StoreId,
		IngredientId:      model.IngredientId,
		IngredientName:    model.IngredientName,
		IngredientGroupId: model.IngredientGroupId,
		Price:             model.Price,
		MarkUp:            model.MarkUp,
		Deleted:           model.deleted,
	}
}

func (s IngredientInfoRepoImpl) CreateIngredient(ctx context.Context, entity *entity.IngredientEntity) error {
	//TODO implement me
	panic("implement me")
}

func (s IngredientInfoRepoImpl) GetIngredient(ctx context.Context, ingredientId string) (*entity.IngredientEntity, error) {
	model := ingredientModel{}
	s.db.Table(s.tableName).Where("ingredient_id = ?", ingredientId).First(&model)
	return model.convertModelToEntity(), nil
}
