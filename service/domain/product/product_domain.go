package product

import (
	"context"
	"e-commerce/service/domain/product/entity"
	"e-commerce/service/domain/product/repo"
	"e-commerce/service/infra/ebus"
	"e-commerce/service/infra/log"
	"e-commerce/service/infra/repo/product"
	"e-commerce/service/infra/utils"
	"errors"
	"fmt"
	"regexp"
	"sync"
)

var once sync.Once

var impl = &ProductDomainImpl{}

type ProductDomainImpl struct {
	SpuRepo      repo.SpuRepo
	SkuRepo      repo.SkuRepo
	CategoryRepo repo.CategoryRepo
}

func NewProductDomainImpl() *ProductDomainImpl {
	impl.SpuRepo = product.NewSpuRepo()
	impl.SkuRepo = product.NewSkuRepo()
	impl.CategoryRepo = product.NewCategoryRepo()
	return impl
}

func (impl *ProductDomainImpl) CreateProduct(ctx context.Context, productName string, categoryId string, skus []entity.SkuEntity, status string, icon string) (*entity.ProductAggInfo, error) {
	// check category
	categoryInfo, err := impl.CategoryRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	fmt.Println("categoryInfo:", categoryInfo)

	spuBO := &entity.SpuEntity{}

	// spu contract
	spuBO, err = impl.spuContract(ctx, spuBO, productName, categoryId, status, icon)
	if err != nil {
		return nil, err
	}

	// create spu
	err = impl.SpuRepo.CreateSpu(ctx, spuBO)
	if err != nil {
		return nil, err
	}

	// skuname list
	skuNames := make([]string, 0)
	// isdefault is only one
	isDefaultTag := false

	for _, sku := range skus {
		if utils.ContainsString(skuNames, sku.SkuName) {
			return nil, errors.New("sku name is duplicate")
		}
		skuNames = append(skuNames, sku.SkuName)

		if isDefaultTag {
			return nil, errors.New("isDefault is only one")
		}
		if sku.IsDefault {
			isDefaultTag = true
		}
	}

	if !isDefaultTag {
		return nil, errors.New("isDefault is not set")
	}

	for _, sku := range skus {
		skuBO, err := impl.skuContract(ctx, &sku, spuBO.SpuId, sku.SkuName, sku.SellAmount, sku.CostAmount, sku.IsDefault, sku.Code)
		if err != nil {
			return nil, err
		}

		err = impl.SkuRepo.CreateSku(ctx, skuBO)
		if err != nil {
			return nil, err
		}
	}

	return entity.ConvertToProductAggInfo(spuBO, skus), nil
}
func (impl *ProductDomainImpl) UpdateProduct(ctx context.Context, spuId string, productName string, categoryId string, skus []entity.SkuEntity, status string, icon string) (*entity.ProductAggInfo, error) {
	// check category
	_, err := impl.CategoryRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	// get spu
	spuBO, err := impl.SpuRepo.GetSpu(ctx, spuId)

	// spu contract
	spuBO, err = impl.spuContract(ctx, spuBO, productName, categoryId, status, icon)
	if err != nil {
		return nil, err
	}

	err = impl.SpuRepo.UpdateSpu(ctx, spuBO)
	if err != nil {
		return nil, err
	}

	// skuname list
	skuNames := make([]string, 0)
	// isdefault is only one
	isDefaultTag := false

	// update sku
	for _, sku := range skus {
		if utils.ContainsString(skuNames, sku.SkuName) {
			return nil, errors.New("sku name is duplicate")
		}
		skuNames = append(skuNames, sku.SkuName)

		if isDefaultTag {
			return nil, errors.New("isDefault is only one")
		}
		if sku.IsDefault {
			isDefaultTag = true
		}
	}

	if !isDefaultTag {
		return nil, errors.New("isDefault is not set")
	}

	for _, sku := range skus {
		skuBO, err := impl.skuContract(ctx, &sku, spuBO.SpuId, sku.SkuName, sku.SellAmount, sku.CostAmount, sku.IsDefault, sku.Code)
		if err != nil {
			return nil, err
		}

		err = impl.SkuRepo.UpdateSku(ctx, skuBO)
		if err != nil {
			return nil, err
		}
	}

	return entity.ConvertToProductAggInfo(spuBO, skus), nil
}
func (impl *ProductDomainImpl) DeleteProduct(ctx context.Context, spuId string) error {

	// delete spu
	err := impl.SpuRepo.DeleteSpu(ctx, spuId)
	if err != nil {
		return err
	}

	// delete sku
	err = impl.SkuRepo.DeleteSku(ctx, spuId)
	if err != nil {
		return err
	}

	// delete inv

	return nil
}
func (impl *ProductDomainImpl) QueryProduct(ctx context.Context, spuId string) (*entity.ProductAggInfo, error) {

	spuBO, err := impl.SpuRepo.GetSpu(ctx, spuId)
	fmt.Println("spuBO:", spuBO)
	if err != nil {
		return nil, err
	}

	skuBOs, err := impl.SkuRepo.GetSkuListBySpuId(ctx, spuId)
	fmt.Println("skuBOs:", skuBOs)
	if err != nil {
		return nil, err
	}

	var skus []entity.SkuEntity
	if skuBOs != nil {
		for _, v := range skuBOs {
			skus = append(skus, *v)
		}
	}

	return &entity.ProductAggInfo{
		SpuId:       spuBO.SpuId,
		CategoryId:  spuBO.CategoryId,
		ProductName: spuBO.ProductName,
		Status:      spuBO.Status,
		Icon:        spuBO.Icon,
		Deleted:     spuBO.Deleted,
		Skus:        skus,
	}, nil
}
func (impl *ProductDomainImpl) QueryProductList(ctx context.Context) ([]*entity.ProductAggInfo, error) {

	spuBOlist, err := impl.SpuRepo.GetSpuList(ctx)
	if err != nil {
		return nil, err
	}

	skuBOList, err := impl.SkuRepo.GetSkuList(ctx)
	if err != nil {
		return nil, err
	}

	productList := make([]*entity.ProductAggInfo, 0)
	skuMap := make(map[string][]entity.SkuEntity)
	for _, v := range skuBOList {
		if _, ok := skuMap[v.SkuName]; !ok {
			skuList := make([]entity.SkuEntity, 0)
			skuList = append(skuList, *v)
			skuMap[v.SpuId] = skuList
		} else {
			skuMap[v.SkuName] = append(skuMap[v.SkuName], *v)
		}
	}

	for _, v := range spuBOlist {
		productList = append(productList, &entity.ProductAggInfo{
			SpuId:       v.SpuId,
			CategoryId:  v.CategoryId,
			ProductName: v.ProductName,
			Status:      v.Status,
			Icon:        v.Icon,
			Deleted:     v.Deleted,
			Skus:        skuMap[v.SpuId],
		})
	}

	return productList, nil
}

func (impl *ProductDomainImpl) spuContract(ctx context.Context, spuBO *entity.SpuEntity, productName string, categoryId string, status string, icon string) (*entity.SpuEntity, error) {
	return spuBO.FillField(ctx, productName, categoryId, status, icon), nil
}

func (impl *ProductDomainImpl) skuContract(ctx context.Context, skuBO *entity.SkuEntity, spuId string, skuName string, sellAmount ebus.Money, costAmount ebus.Money, isDefault bool, code string) (*entity.SkuEntity, error) {
	re := regexp.MustCompile("^[a-zA-Z0-9]+$")
	if !re.MatchString(code) {
		log.Error("code is not allowed")
	}

	return skuBO.FillField(ctx, spuId, skuName, sellAmount, costAmount, isDefault, code), nil
}
