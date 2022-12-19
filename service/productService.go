package service

import (
	"github.com/highergear/go-ecommerence/model"
	"github.com/highergear/go-ecommerence/model/input"
)

type ProductService interface {
	Create(i input.CreateProductInput, uid uint) (model.Product, error)
	GetProductList(limit int, offset int) []model.Product
	GetProductById(id uint) model.Product
	GetProductListBySellerId(sellerId uint, limit int, offset int) []model.Product
}
