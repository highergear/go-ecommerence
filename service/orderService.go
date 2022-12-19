package service

import (
	"github.com/highergear/go-ecommerence/model"
	"github.com/highergear/go-ecommerence/model/input"
)

type OrderService interface {
	Create(i input.CreateOrderInput, product model.Product, uid uint) (model.Order, error)
	GetOrderById(id uint) model.Order
	GetOrderListByBuyerId(buyerId uint, limit int, offset int) []model.Order
	GetOrderListBySellerId(sellerId uint, limit int, offset int) []model.Order
	UpdateStatusToAccepted(order model.Order) (model.Order, error)
}
