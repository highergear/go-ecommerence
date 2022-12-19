package repository

import "github.com/highergear/go-ecommerence/model"

type ProductRepository interface {
	Save(product model.Product) (model.Product, error)
	FindAll(limit int, offset int) []model.Product
	FindById(id uint) model.Product
	FindBySellerId(sellerId uint, limit int, offset int) []model.Product
}
