package entity

type ProductAggInfo struct {
	// spu info
	SpuId       string
	CategoryId  string
	ProductName string
	Status      string
	Icon        string
	Deleted     bool

	// sku info
	Skus []SkuEntity
}
