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

var CustomizationTableName = "customization"

var CustomizationTableNameRepoImpl = &CustomizationInfoRepoImpl{}

type CustomizationInfoRepoImpl struct {
	db        *gorm.DB
	tableName string
}

func NewCustomizationRepo() repo.CustomizationRepo {
	return &CustomizationInfoRepoImpl{
		db:        orm.GetORM(consts.DB_NAME),
		tableName: CustomizationTableName,
	}
}

type customizationModel struct {
	StoreId              string     `gorm:"column:store_id"`
	CustomizationId      string     `gorm:"column:customization_id"`
	CustomizationName    string     `gorm:"column:customization_name"`
	Deleted              bool       `gorm:"column:deleted"`
	CustomizationGroupId string     `gorm:"column:customization_group_id"`
	Price                ebus.Money `gorm:"column:price"`
	MarkUp               bool       `gorm:"column:mark_up"`
	deleted              bool       `gorm:"column:deleted"`
}

func (model customizationModel) convertEntityToModel(entity *entity.CustomizationEntity) {
	model.StoreId = entity.StoreId
	model.CustomizationId = entity.CustomizationId
	model.CustomizationName = entity.CustomizationName
	model.CustomizationGroupId = entity.CustomizationGroupId
	model.Price = entity.Price
	model.MarkUp = entity.MarkUp
	model.deleted = entity.Deleted
}

// model转entity
func (model customizationModel) convertModelToEntity() *entity.CustomizationEntity {
	return &entity.CustomizationEntity{
		StoreId:              model.StoreId,
		CustomizationId:      model.CustomizationId,
		CustomizationName:    model.CustomizationName,
		CustomizationGroupId: model.CustomizationGroupId,
		Price:                model.Price,
		MarkUp:               model.MarkUp,
		Deleted:              model.deleted,
	}
}

func (s CustomizationInfoRepoImpl) CreateCustomization(ctx context.Context, customization *entity.CustomizationEntity) error {
	//TODO implement me
	panic("implement me")
}

func (s CustomizationInfoRepoImpl) GetCustomization(ctx context.Context, customizationId string) (*entity.CustomizationEntity, error) {
	model := customizationModel{}
	s.db.Table(s.tableName).Where("customization_id = ?", customizationId).First(&model)
	return model.convertModelToEntity(), nil
}
