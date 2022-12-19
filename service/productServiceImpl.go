package service

import (
	"github.com/highergear/go-ecommerence/model"
	"github.com/highergear/go-ecommerence/model/input"
	"github.com/highergear/go-ecommerence/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(ProductRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		ProductRepository: ProductRepository,
	}
}

func (service *ProductServiceImpl) Create(i input.CreateProductInput, uid uint) (model.Product, error) {
	product := model.Product{}
	product.Name = i.Name
	product.Description = i.Description
	product.Price = i.Price
	product.SellerID = uid

	savedProduct, err := service.ProductRepository.Save(product)
	if err != nil {
		return model.Product{}, err
	}
	return savedProduct, nil
}

func (service *ProductServiceImpl) GetProductList(limit int, offset int) []model.Product {
	return service.ProductRepository.FindAll(limit, offset)
}

func (service *ProductServiceImpl) GetProductById(id uint) model.Product {
	return service.ProductRepository.FindById(id)
}

func (service *ProductServiceImpl) GetProductListBySellerId(sellerId uint, limit int, offset int) []model.Product {
	return service.ProductRepository.FindBySellerId(sellerId, limit, offset)
}
