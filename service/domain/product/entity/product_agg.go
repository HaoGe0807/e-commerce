package entity

import (
	"e-commerce/service/infra/log"
	"encoding/json"
)

type ProductAggInfo struct {
	// spu info
	SpuId       string `json:"spu_id"`
	CategoryId  string `json:"category_id"`
	ProductName string `json:"product_name"`
	Status      string `json:"status"`
	Icon        string `json:"icon"`
	Deleted     bool   `json:"deleted"`

	// sku info
	Skus []SkuEntity `json:"skus"`
}

func ConvertToProductAggInfo(entity *SpuEntity, skus []*SkuEntity) *ProductAggInfo {
	return &ProductAggInfo{
		SpuId:       entity.SpuId,
		CategoryId:  entity.CategoryId,
		ProductName: entity.ProductName,
		Status:      entity.Status,
		Icon:        entity.Icon,
		Deleted:     entity.Deleted,
		Skus:        ConvertSkuAddressToValue(skus),
	}
}

func ConvertSkuAddressToValue(skus []*SkuEntity) []SkuEntity {
	var skusEntity []SkuEntity
	for _, sku := range skus {
		skusEntity = append(skusEntity, *sku)
	}
	return skusEntity
}

func ConvertStringToProductAggInfo(str string) (*ProductAggInfo, error) {
	productAggInfo := &ProductAggInfo{}
	err := json.Unmarshal([]byte(str), productAggInfo)
	if err != nil {
		log.Error("Error converting Redis string to Spu struct: %v\n", err)
		log.Error("The Redis string content is: %s\n", str)
		return &ProductAggInfo{}, err
	}

	return productAggInfo, nil
}

func ProductAggInfoToJsonMarshal(productAggInfo *ProductAggInfo) []byte {
	marshal, err := json.Marshal(productAggInfo)
	if err != nil {
		return nil
	}
	return marshal
}
