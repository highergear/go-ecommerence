package repository

import "github.com/highergear/go-ecommerence/model"

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository ProductRepositoryImpl) Save(product model.Product) (model.Product, error) {
	err := DB.Create(&product).Error
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (repository ProductRepositoryImpl) FindAll(limit int, offset int) []model.Product {
	var productList []model.Product
	err := DB.Model(model.Product{}).Limit(limit).Offset(offset).Find(&productList).Error
	if err != nil {
		return []model.Product{}
	}
	return productList
}

func (repository ProductRepositoryImpl) FindById(id uint) model.Product {
	var product model.Product
	err := DB.Model(model.Product{}).Where("id = ?", id).Take(&product).Error
	if err != nil {
		return model.Product{}
	}
	return product
}

func (repository ProductRepositoryImpl) FindBySellerId(sellerId uint, limit int, offset int) []model.Product {
	var productList []model.Product
	err := DB.Model(model.Product{}).Where("seller_id = ?", sellerId).Limit(limit).Offset(offset).Find(&productList).Error
	if err != nil {
		return []model.Product{}
	}
	return productList
}
