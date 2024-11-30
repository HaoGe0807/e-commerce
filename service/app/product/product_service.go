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
	CreateProduct(ctx context.Context, productName string, categoryId string, skus []entity.SkuEntity, unit string, mnemonicCode string, status string, storeId string, customizationList []entity.CustomizationEntity, ingredientList []entity.IngredientEntity, icon string, priceMethod string, shape string, shapeColor string, firstDisplay string, productSpecifications string, productType string) error
	UpdateProduct(ctx context.Context, spuId string, productName string, categoryId string, skus []entity.SkuEntity, unit string, mnemonicCode string, status string, storeId string, customizationList []entity.CustomizationEntity, ingredientList []entity.IngredientEntity, icon string, priceMethod string, shape string, shapeColor string, firstDisplay string, productSpecifications string, productType string) error
	DeleteProduct(ctx context.Context, storeId, spuId string) error
	QueryProduct(ctx context.Context, storeId, spuId string) (productVO.ProductVO, error)
	QueryProductList(ctx context.Context, storeId string) ([]productVO.ProductVO, error)
}

func (p ProductService) CreateProduct(ctx context.Context, productName string, categoryId string, skus []entity.SkuEntity, unit string, mnemonicCode string, status string, storeId string, customizationList []entity.CustomizationEntity, ingredientList []entity.IngredientEntity, icon string, priceMethod string, shape string, shapeColor string, firstDisplay string, productSpecifications string, productType string) error {
	// create product
	_, err := p.impl.CreateProduct(ctx, storeId, productName, categoryId, skus, unit, mnemonicCode, status, icon, priceMethod, shape, shapeColor, firstDisplay, productSpecifications, productType)
	if err != nil {
		return err
	}

	// create inv

	return nil
}

func (p ProductService) UpdateProduct(ctx context.Context, spuId string, productName string, categoryId string, skus []entity.SkuEntity, unit string, mnemonicCode string, status string, storeId string, customizationList []entity.CustomizationEntity, ingredientList []entity.IngredientEntity, icon string, priceMethod string, shape string, shapeColor string, firstDisplay string, productSpecifications string, productType string) error {
	// update product
	err := p.impl.UpdateProduct(ctx, spuId, productName, categoryId, skus, unit, mnemonicCode, status, storeId, customizationList, ingredientList, icon, priceMethod, shape, shapeColor, firstDisplay, productSpecifications, productType)
	if err != nil {
		return err
	}

	return nil
}

func (p ProductService) DeleteProduct(ctx context.Context, storeId, spuId string) error {
	err := p.impl.DeleteProduct(ctx, storeId, spuId)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductService) QueryProduct(ctx context.Context, storeId, spuId string) (productVO.ProductVO, error) {

	productInfo, err := p.impl.QueryProduct(ctx, storeId, spuId)
	if err != nil {
		return productVO.ProductVO{}, err
	}

	// 查询库存
	inv := []string{}

	return productVO.ProductBOToVO(productInfo, inv), nil
}

func (p ProductService) QueryProductList(ctx context.Context, storeId string) ([]productVO.ProductVO, error) {
	productList, err := p.impl.QueryProductList(ctx, storeId)
	if err != nil {
		return nil, err
	}

	// 查询库存
	inv := []string{}
	resp := make([]productVO.ProductVO, len(productList))

	for _, info := range productList {
		resp = append(resp, productVO.ProductBOToVO(info, inv))
	}
	return resp, nil
}
