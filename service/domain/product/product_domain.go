package product

import (
	"context"
	"e-commerce/service/domain/product/entity"
	"e-commerce/service/domain/product/repo"
	"e-commerce/service/infra/consts"
	"e-commerce/service/infra/ebus"
	"e-commerce/service/infra/errors"
	"e-commerce/service/infra/log"
	"e-commerce/service/infra/redis"
	"e-commerce/service/infra/repo/product"
	"e-commerce/service/infra/utils"
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

	if categoryInfo.CategoryId == "" {
		return nil, errors.ErrorEnum(errors.CATEGORY_NOT_EXIST)
	}

	spuBO := &entity.SpuEntity{}
	skuBOList := make([]*entity.SkuEntity, 0)

	// spu contract
	spuBO, err = impl.spuContract(ctx, spuBO, productName, categoryId, status, icon)
	if err != nil {
		return nil, err
	}

	// skuname list
	skuNames := make([]string, 0)
	// isdefault is only one
	isDefaultTag := false

	for _, sku := range skus {
		if utils.ContainsString(skuNames, sku.SkuName) {
			return nil, errors.ErrorEnum(errors.SKU_NAME_DUPLICATE)
		}
		skuNames = append(skuNames, sku.SkuName)

		if isDefaultTag && sku.IsDefault {
			return nil, errors.ErrorEnum(errors.PRODUCT_ONLY_ONE_DEFAULT_SKU)
		}
		if sku.IsDefault {
			isDefaultTag = true
		}
	}

	if !isDefaultTag {
		return nil, errors.ErrorEnum(errors.PRODUCT_ONLY_ONE_DEFAULT_SKU)
	}

	for _, sku := range skus {
		skuBO := &entity.SkuEntity{}
		skuBO, err := impl.skuContract(ctx, &sku, spuBO.SpuId, sku.SkuName, sku.SellAmount, sku.CostAmount, sku.IsDefault, sku.Code)
		if err != nil {
			return nil, err
		}

		skuBOList = append(skuBOList, skuBO)
	}

	// create spu
	err = impl.SpuRepo.CreateSpu(ctx, spuBO)
	if err != nil {
		return nil, err
	}
	err = impl.SkuRepo.SaveSkuListBySpuId(ctx, skuBOList, spuBO.SpuId)
	if err != nil {
		return nil, err
	}

	productAggInfo := entity.ConvertToProductAggInfo(spuBO, skuBOList)
	redis.GetECommerceRedis().SetNX(consts.REDIS_PRODUCT+spuBO.SpuId, entity.ProductAggInfoToJsonMarshal(productAggInfo), consts.REDIS_DEFAULT_TIME)
	return productAggInfo, nil
}
func (impl *ProductDomainImpl) UpdateProduct(ctx context.Context, spuId string, productName string, categoryId string, skus []entity.SkuEntity, status string, icon string) (*entity.ProductAggInfo, error) {
	// delete redis
	redis.GetECommerceRedis().Del(consts.REDIS_PRODUCT + spuId)

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

	// skuname list
	skuNames := make([]string, 0)
	// isdefault is only one
	isDefaultTag := false

	// update sku
	for _, sku := range skus {
		if utils.ContainsString(skuNames, sku.SkuName) {
			return nil, errors.ErrorEnum(errors.SKU_NAME_DUPLICATE)
		}
		skuNames = append(skuNames, sku.SkuName)

		if isDefaultTag {
			return nil, errors.ErrorEnum(errors.PRODUCT_ONLY_ONE_DEFAULT_SKU)
		}
		if sku.IsDefault {
			isDefaultTag = true
		}
	}

	if !isDefaultTag {
		return nil, errors.ErrorEnum(errors.PRODUCT_ONLY_ONE_DEFAULT_SKU)
	}

	skuBOList := make([]*entity.SkuEntity, 0)
	for _, sku := range skus {
		skuBO := &entity.SkuEntity{}
		skuBO, err := impl.skuContract(ctx, &sku, spuBO.SpuId, sku.SkuName, sku.SellAmount, sku.CostAmount, sku.IsDefault, sku.Code)
		if err != nil {
			return nil, err
		}

		skuBOList = append(skuBOList, skuBO)
	}

	err = impl.SpuRepo.UpdateSpu(ctx, spuBO)
	if err != nil {
		return nil, err
	}

	err = impl.SkuRepo.SaveSkuListBySpuId(ctx, skuBOList, spuBO.SpuId)
	if err != nil {
		return nil, err
	}

	productAggInfo := entity.ConvertToProductAggInfo(spuBO, skuBOList)
	redis.GetECommerceRedis().SetNX(consts.REDIS_PRODUCT+spuBO.SpuId, entity.ProductAggInfoToJsonMarshal(productAggInfo), consts.REDIS_DEFAULT_TIME)
	return productAggInfo, nil
}
func (impl *ProductDomainImpl) DeleteProduct(ctx context.Context, spuId string) error {
	// delete redis
	redis.GetECommerceRedis().Del(consts.REDIS_PRODUCT + spuId)

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

	return nil
}
func (impl *ProductDomainImpl) QueryProduct(ctx context.Context, spuId string) (*entity.ProductAggInfo, error) {
	redisValue := redis.GetECommerceRedis().Get(consts.REDIS_PRODUCT + spuId).Val()
	if redisValue != "" {
		log.Infof("QueryProduct redisValue:%s", redisValue)
		return entity.ConvertStringToProductAggInfo(redisValue)
	}

	spuBO, err := impl.SpuRepo.GetSpu(ctx, spuId)
	if err != nil {
		return nil, err
	}

	skuBOs, err := impl.SkuRepo.GetSkuListBySpuId(ctx, spuId)
	if err != nil {
		return nil, err
	}

	return &entity.ProductAggInfo{
		SpuId:       spuBO.SpuId,
		CategoryId:  spuBO.CategoryId,
		ProductName: spuBO.ProductName,
		Status:      spuBO.Status,
		Icon:        spuBO.Icon,
		Deleted:     spuBO.Deleted,
		Skus:        entity.ConvertSkuAddressToValue(skuBOs),
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
	skuMap := make(map[string][]*entity.SkuEntity)
	for _, v := range skuBOList {
		if _, ok := skuMap[v.SpuId]; !ok {
			skuList := make([]*entity.SkuEntity, 0)
			skuList = append(skuList, v)
			skuMap[v.SpuId] = skuList
		} else {
			skuMap[v.SpuId] = append(skuMap[v.SpuId], v)
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
			Skus:        entity.ConvertSkuAddressToValue(skuMap[v.SpuId]),
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
		return nil, errors.ErrorEnum(errors.SKU_CODE_ERROR)
	}

	return skuBO.FillField(ctx, spuId, skuName, sellAmount, costAmount, isDefault, code), nil
}

// category
// createCategory
func (impl *ProductDomainImpl) CreateCategory(ctx context.Context, categoryName string) (*entity.CategoryEntity, error) {
	// verify name
	categorySameName, err := impl.CategoryRepo.GetCategoryByName(ctx, "", categoryName)
	if err != nil {
		return nil, err
	}
	if categorySameName.CategoryId != "" {
		return nil, errors.ErrorEnum(errors.CATEGORY_NAME_EXIST)
	}

	categoryBO := &entity.CategoryEntity{}
	categoryBO = categoryBO.FillField(ctx, categoryName)

	err = impl.CategoryRepo.CreateCategory(ctx, categoryBO)
	if err != nil {
		return nil, err
	}
	fmt.Println("categoryBO", categoryBO)
	redis.GetECommerceRedis().SetNX(consts.REDIS_CATEGOEY+categoryBO.CategoryId, entity.CategoryInfoToJsonMarshal(categoryBO), consts.REDIS_DEFAULT_TIME)
	return categoryBO, nil
}

func (impl *ProductDomainImpl) UpdateCategory(ctx context.Context, categoryId, categoryName string) (*entity.CategoryEntity, error) {
	redis.GetECommerceRedis().Del(consts.REDIS_CATEGOEY + categoryId)

	// get category
	categoryBO, err := impl.CategoryRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	// verify name
	categorySameName, err := impl.CategoryRepo.GetCategoryByName(ctx, categoryId, categoryName)
	if err != nil {
		return nil, err
	}
	if categorySameName.CategoryId != "" {
		return nil, errors.ErrorEnum(errors.CATEGORY_NAME_EXIST)
	}

	categoryBO = categoryBO.FillField(ctx, categoryName)

	err = impl.CategoryRepo.UpdateCategory(ctx, categoryBO)
	if err != nil {
		return nil, err
	}

	redis.GetECommerceRedis().SetNX(consts.REDIS_CATEGOEY+categoryBO.CategoryId, entity.CategoryInfoToJsonMarshal(categoryBO), consts.REDIS_DEFAULT_TIME)
	return categoryBO, nil
}

func (impl *ProductDomainImpl) DeleteCategory(ctx context.Context, categoryId string) error {
	redis.GetECommerceRedis().Del(consts.REDIS_CATEGOEY + categoryId)

	// get category
	categoryBO, err := impl.CategoryRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return err
	}
	if categoryBO.CategoryId != "" {
		return errors.ErrorEnum(errors.CATEGORY_NAME_EXIST)
	}

	// verify product
	spuBOList, err := impl.SpuRepo.GetSpuListByCategoryId(ctx, categoryId)
	if err != nil {
		return err
	}

	if len(spuBOList) > 0 {
		return errors.ErrorEnum(errors.CATEGORY_EXIST_PRODUCT)
	}

	err = impl.CategoryRepo.DeleteCategory(ctx, categoryId)
	if err != nil {
		return err
	}

	return nil
}

func (impl *ProductDomainImpl) QueryCategory(ctx context.Context, categoryId string) (*entity.CategoryEntity, error) {
	redisValue := redis.GetECommerceRedis().Get(consts.REDIS_CATEGOEY + categoryId).Val()
	if redisValue != "" {
		log.Infof("QueryCategory redisValue:%s", redisValue)
		return entity.ConvertStringToCategoryInfo(redisValue)
	}

	// get category
	categoryBO, err := impl.CategoryRepo.GetCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	return categoryBO, nil
}

func (impl *ProductDomainImpl) QueryCategoryList(ctx context.Context) ([]*entity.CategoryEntity, error) {
	// get category
	categoryBOList, err := impl.CategoryRepo.GetCategoryList(ctx)
	if err != nil {
		return nil, err
	}

	return categoryBOList, nil
}
