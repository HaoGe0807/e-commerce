package entity

type ProductAggInfo struct {
	// spu info
	SpuId                 string
	CategoryId            string
	StoreId               string
	ProductName           string
	Unit                  string
	Status                string
	MnemonicCode          string
	ProductSpecifications string
	Icon                  string
	CustomizationList     string
	IngredientList        string
	Deleted               bool
	PriceMethod           string
	Sort                  int
	SortFiled             string
	Shape                 string
	ShapeColor            string
	ProductType           string
	Version               string
	FirstDisplay          string

	// sku info
	Skus []SkuEntity
}
