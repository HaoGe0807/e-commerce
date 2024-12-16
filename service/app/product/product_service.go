package product

import (
	"context"
	productVO "e-commerce/service/app/product/vo"
	"e-commerce/service/domain/product"
	"e-commerce/service/domain/product/entity"
)

type ProductService struct {
	impl *product.ProductDomainImpl
}

func NewProductService() PdService {
	PdSvc := &ProductService{
		impl: product.NewProductDomainImpl(),
	}
	return PdSvc
}

type PdService interface {
	CreateProduct(ctx context.Context, productName string, categoryId string, skus []entity.SkuEntity, status string, icon string) (*productVO.ProductVO, error)
	UpdateProduct(ctx context.Context, spuId string, productName string, categoryId string, skus []entity.SkuEntity, status string, icon string) (*productVO.ProductVO, error)
	DeleteProduct(ctx context.Context, spuId string) error
	QueryProduct(ctx context.Context, spuId string) (*productVO.ProductVO, error)
	QueryProductList(ctx context.Context) ([]*productVO.ProductVO, error)
}

func (p ProductService) CreateProduct(ctx context.Context, productName string, categoryId string, skus []entity.SkuEntity, status string, icon string) (*productVO.ProductVO, error) {
	// create product
	productAggInfo, err := p.impl.CreateProduct(ctx, productName, categoryId, skus, status, icon)
	if err != nil {
		return nil, err
	}

	return productVO.ProductBOToVO(productAggInfo), nil
}

func (p ProductService) UpdateProduct(ctx context.Context, spuId string, productName string, categoryId string, skus []entity.SkuEntity, status string, icon string) (*productVO.ProductVO, error) {
	// update product
	productAggInfo, err := p.impl.UpdateProduct(ctx, spuId, productName, categoryId, skus, status, icon)
	if err != nil {
		return nil, err
	}

	return productVO.ProductBOToVO(productAggInfo), nil
}

func (p ProductService) DeleteProduct(ctx context.Context, spuId string) error {
	err := p.impl.DeleteProduct(ctx, spuId)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductService) QueryProduct(ctx context.Context, spuId string) (*productVO.ProductVO, error) {

	productInfo, err := p.impl.QueryProduct(ctx, spuId)
	if err != nil {
		return nil, err
	}

	return productVO.ProductBOToVO(productInfo), nil
}

func (p ProductService) QueryProductList(ctx context.Context) ([]*productVO.ProductVO, error) {
	productList, err := p.impl.QueryProductList(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*productVO.ProductVO, 0)

	for _, info := range productList {
		resp = append(resp, productVO.ProductBOToVO(info))
	}
	return resp, nil
}
