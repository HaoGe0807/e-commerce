package product

import (
	"context"
	"e-commerce/service/domain/product/entity"
	"e-commerce/service/domain/product/repo"
	"e-commerce/service/infra/consts"
	"e-commerce/service/infra/ebus"
	"e-commerce/service/infra/log"
	"e-commerce/service/infra/repo/product"
	"e-commerce/service/infra/utils"
	"errors"
	"regexp"
	"sync"
)

var once sync.Once

var impl = &ProductDomainImpl{}

type ProductDomainImpl struct {
	SpuRepo           repo.SpuRepo
	SkuRepo           repo.SkuRepo
	CategoryRepo      repo.CategoryRepo
	CustomizationRepo repo.CustomizationRepo
	IngredientRepo    repo.IngredientRepo
}

func NewProductDomainImpl() *ProductDomainImpl {
	impl.SpuRepo = product.NewSpuRepo()
	impl.SkuRepo = product.NewSkuRepo()
	impl.CategoryRepo = product.NewCategoryRepo()
	impl.CustomizationRepo = product.NewCustomizationRepo()
	impl.IngredientRepo = product.NewIngredientRepo()
	return impl
}

func (impl *ProductDomainImpl) CreateProduct(ctx context.Context, storeId string, productName string, categoryId string, skus []entity.SkuEntity, unit string, mnemonicCode string, status string, icon string, priceMethod string, shape string, shapeColor string, firstDisplay string, productSpecifications string, productType string) (string, error) {

	// check category
	_, err := impl.CategoryRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return "", err
	}

	spuBO := &entity.SpuEntity{}

	// spu contract
	spuBO, err = impl.spuContract(ctx, spuBO, storeId, productName, unit, categoryId, mnemonicCode, status, icon, priceMethod, productSpecifications, shape, shapeColor, firstDisplay, productType)
	if err != nil {
		return "", err
	}

	// create spu
	err = impl.SpuRepo.CreateSpu(ctx, spuBO)
	if err != nil {
		return "", err
	}

	// create sku
	if productSpecifications == consts.SINGLE && len(skus) != 1 {
		return "", errors.New("single product must have one sku")
	}

	// skuname list
	skuNames := make([]string, 0)
	// isdefault is only one
	isDefaultTag := false

	for _, sku := range skus {
		if utils.ContainsString(skuNames, sku.SkuName) {
			return "", errors.New("sku name is duplicate")
		}
		skuNames = append(skuNames, sku.SkuName)

		if isDefaultTag {
			return "", errors.New("isDefault is only one")
		}
		if sku.IsDefault {
			isDefaultTag = true
		}
	}

	if !isDefaultTag {
		return "", errors.New("isDefault is not set")
	}

	for _, sku := range skus {
		skuBO, err := impl.skuContract(ctx, &sku, storeId, spuBO.SpuId, sku.SkuName, sku.SellAmount, sku.CostAmount, sku.IsDefault, sku.Code, sku.MinimumStock)
		if err != nil {
			return "", err
		}

		err = impl.SkuRepo.CreateSku(ctx, skuBO)
		if err != nil {
			return "", err
		}
	}

	return spuBO.SpuId, nil
}
func (impl *ProductDomainImpl) UpdateProduct(ctx context.Context, spuId string, productName string, categoryId string, skus []entity.SkuEntity, unit string, mnemonicCode string, status string, storeId string, customizationList []entity.CustomizationEntity, ingredientList []entity.IngredientEntity, icon string, priceMethod string, shape string, shapeColor string, firstDisplay string, productSpecifications string, productType string) error {
	// check category
	_, err := impl.CategoryRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return err
	}

	// get spu
	spuBO, err := impl.SpuRepo.GetSpu(ctx, storeId, spuId)

	// spu contract
	spuBO, err = impl.spuContract(ctx, spuBO, storeId, productName, unit, categoryId, mnemonicCode, status, icon, priceMethod, productSpecifications, shape, shapeColor, firstDisplay, productType)
	if err != nil {
		return err
	}

	err = impl.SpuRepo.UpdateSpu(ctx, spuBO)
	if err != nil {
		return err
	}

	// skuname list
	skuNames := make([]string, 0)
	// isdefault is only one
	isDefaultTag := false

	// update sku
	for _, sku := range skus {
		if utils.ContainsString(skuNames, sku.SkuName) {
			return errors.New("sku name is duplicate")
		}
		skuNames = append(skuNames, sku.SkuName)

		if isDefaultTag {
			return errors.New("isDefault is only one")
		}
		if sku.IsDefault {
			isDefaultTag = true
		}
	}

	if !isDefaultTag {
		return errors.New("isDefault is not set")
	}

	for _, sku := range skus {
		skuBO, err := impl.skuContract(ctx, &sku, storeId, spuBO.SpuId, sku.SkuName, sku.SellAmount, sku.CostAmount, sku.IsDefault, sku.Code, sku.MinimumStock)
		if err != nil {
			return err
		}

		err = impl.SkuRepo.UpdateSku(ctx, skuBO)
		if err != nil {
			return err
		}
	}

	return nil
}
func (impl *ProductDomainImpl) DeleteProduct(ctx context.Context, storeId, spuId string) error {

	// delete spu
	err := impl.SpuRepo.DeleteSpu(ctx, storeId, spuId)
	if err != nil {
		return err
	}

	// delete sku
	err = impl.SkuRepo.DeleteSku(ctx, storeId, spuId)
	if err != nil {
		return err
	}

	// delete inv

	return nil
}
func (impl *ProductDomainImpl) QueryProduct(ctx context.Context, storeId, spuId string) (entity.ProductAggInfo, error) {

	spuBO, err := impl.SpuRepo.GetSpu(ctx, storeId, spuId)
	if err != nil {
		return entity.ProductAggInfo{}, err
	}

	skuBOs, err := impl.SkuRepo.GetSkuListByStoreIdAndSpuId(ctx, storeId, spuId)
	if err != nil {
		return entity.ProductAggInfo{}, err
	}

	var skus []entity.SkuEntity
	if skuBOs != nil {
		for _, v := range skuBOs {
			skus = append(skus, *v)
		}
	}

	return entity.ProductAggInfo{
		SpuId:                 spuBO.SpuId,
		CategoryId:            spuBO.CategoryId,
		StoreId:               spuBO.StoreId,
		ProductName:           spuBO.ProductName,
		Unit:                  spuBO.Unit,
		Status:                spuBO.Status,
		MnemonicCode:          spuBO.MnemonicCode,
		ProductSpecifications: spuBO.ProductSpecifications,
		Icon:                  spuBO.Icon,
		Deleted:               spuBO.Deleted,
		PriceMethod:           spuBO.PriceMethod,
		Sort:                  spuBO.Sort,
		SortFiled:             spuBO.SortFiled,
		Shape:                 spuBO.Shape,
		ShapeColor:            spuBO.ShapeColor,
		ProductType:           spuBO.ProductType,
		Version:               spuBO.Version,
		FirstDisplay:          spuBO.FirstDisplay,
		Skus:                  skus,
	}, nil
}
func (impl *ProductDomainImpl) QueryProductList(ctx context.Context, storeId string) ([]entity.ProductAggInfo, error) {

	spuBOlist, err := impl.SpuRepo.GetSpuListByStoreId(ctx, storeId)
	if err != nil {
		return nil, err
	}

	skuBOList, err := impl.SkuRepo.GetSkuListByStoreId(ctx, storeId)
	if err != nil {
		return nil, err
	}

	productList := make([]entity.ProductAggInfo, len(spuBOlist))
	skuMap := make(map[string][]entity.SkuEntity)
	for _, v := range skuBOList {
		if _, ok := skuMap[v.SkuName]; !ok {
			skuList := make([]entity.SkuEntity, 0)
			skuList = append(skuList, *v)
			skuMap[v.SpuId] = skuList
		}
		skuMap[v.SkuName] = append(skuMap[v.SkuName], *v)
	}

	for _, v := range spuBOlist {
		productList = append(productList, entity.ProductAggInfo{
			SpuId:                 v.SpuId,
			CategoryId:            v.CategoryId,
			StoreId:               v.StoreId,
			ProductName:           v.ProductName,
			Unit:                  v.Unit,
			Status:                v.Status,
			MnemonicCode:          v.MnemonicCode,
			ProductSpecifications: v.ProductSpecifications,
			Icon:                  v.Icon,
			CustomizationList:     v.CustomizationList,
			IngredientList:        v.IngredientList,
			Deleted:               v.Deleted,
			PriceMethod:           v.PriceMethod,
			Sort:                  v.Sort,
			SortFiled:             v.SortFiled,
			Shape:                 v.Shape,
			ShapeColor:            v.ShapeColor,
			ProductType:           v.ProductType,
			Version:               v.Version,
			FirstDisplay:          v.FirstDisplay,
			Skus:                  skuMap[v.SpuId],
		})
	}

	return productList, nil
}

func (impl *ProductDomainImpl) spuContract(ctx context.Context, spuBO *entity.SpuEntity, storeId string, productName string, unit string, categoryId string, mnemonicCode string, status string, icon string, priceMethod string, productSpecifications string, shape string, shapeColor string, firstDisplay string, productType string) (*entity.SpuEntity, error) {
	if !utils.ContainsString(consts.ALLOWSHAPES, shape) {
		log.Error("shape is not allowed")
		return nil, errors.New("shape is not allowed")
	}

	return spuBO.FillField(ctx, storeId, productName, unit, categoryId, mnemonicCode, status, icon, priceMethod, productSpecifications, shape, shapeColor, firstDisplay, productType), nil
}

func (impl *ProductDomainImpl) skuContract(ctx context.Context, skuBO *entity.SkuEntity, storeId string, spuId string, skuName string, sellAmount ebus.Money, costAmount ebus.Money, isDefault bool, code string, minimumStock int64) (*entity.SkuEntity, error) {
	re := regexp.MustCompile("^[a-zA-Z0-9]+$")
	if !re.MatchString(code) {
		log.Error("code is not allowed")
	}

	return skuBO.FillField(ctx, storeId, spuId, skuName, sellAmount, costAmount, isDefault, code, minimumStock), nil
}
