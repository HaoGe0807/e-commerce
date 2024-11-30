package app

import (
	"e-commerce/service/app/product"
)

type IRetailService interface {
	product.PdService
}

type RetailService struct {
	product.PdService
}

func NewService() IRetailService {
	svc := &RetailService{
		PdService: product.NewProductService(),
	}
	return svc
}
