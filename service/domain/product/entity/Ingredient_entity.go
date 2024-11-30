package entity

import "e-commerce/service/infra/ebus"

type IngredientEntity struct {
	StoreId           string
	IngredientId      string
	IngredientName    string
	Deleted           bool
	IngredientGroupId string
	Price             ebus.Money
	MarkUp            bool
}
