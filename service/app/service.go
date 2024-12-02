package app

import (
	"e-commerce/service/app/product"
)

type ECommerceService interface {
	product.PdService
}

type CommerceService struct {
	product.PdService
}

func NewService() ECommerceService {
	svc := &CommerceService{
		PdService: product.NewProductService(),
	}
	return svc
}
