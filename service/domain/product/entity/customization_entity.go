package entity

import "e-commerce/service/infra/ebus"

type CustomizationEntity struct {
	StoreId              string
	CustomizationId      string
	CustomizationName    string
	CustomizationGroupId string
	Price                ebus.Money
	MarkUp               bool
	Deleted              bool
}
