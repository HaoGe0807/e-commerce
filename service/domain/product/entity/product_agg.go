package entity

type ProductAggInfo struct {
	// spu info
	SpuId       string `json:"spu_id"`
	CategoryId  string `json:"category_id"`
	ProductName string `json:"product_name"`
	Status      string `json:"status"`
	Icon        string `json:"icon"`
	Deleted     bool   `json:"deleted"`

	// sku info
	Skus []*SkuEntity `json:"skus"`
}

func ConvertToProductAggInfo(entity *SpuEntity, skus []*SkuEntity) *ProductAggInfo {
	return &ProductAggInfo{
		SpuId:       entity.SpuId,
		CategoryId:  entity.CategoryId,
		ProductName: entity.ProductName,
		Status:      entity.Status,
		Icon:        entity.Icon,
		Deleted:     entity.Deleted,
		Skus:        skus,
	}
}
