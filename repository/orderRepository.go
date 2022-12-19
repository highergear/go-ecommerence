package repository

import "github.com/highergear/go-ecommerence/model"

type OrderRepository interface {
	Save(order model.Order) (model.Order, error)
	FindById(id uint) model.Order
	FindBySellerId(sellerId uint, limit int, offset int) []model.Order
	FindByBuyerId(buyerId uint, limit int, offset int) []model.Order
	UpdateStatusToAccepted(order model.Order) (model.Order, error)
}
