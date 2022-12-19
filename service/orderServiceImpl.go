package service

import (
	"github.com/highergear/go-ecommerence/model"
	"github.com/highergear/go-ecommerence/model/input"
	"github.com/highergear/go-ecommerence/repository"
	"github.com/highergear/go-ecommerence/utils"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
	}
}

func (service *OrderServiceImpl) Create(i input.CreateOrderInput, product model.Product, uid uint) (model.Order, error) {
	order := MapInputToOrder(i, product, uid)
	savedOrder, err := service.OrderRepository.Save(order)
	if err != nil {
		return model.Order{}, err
	}
	return savedOrder, nil
}

func MapInputToOrder(i input.CreateOrderInput, product model.Product, uid uint) model.Order {
	order := model.Order{}
	order.DeliverySourceAddress = i.DeliverySourceAddress
	order.DeliveryDestAddress = i.DeliveryDestAddress
	order.Items = i.Items
	order.Quantity = i.Quantity
	order.Status = utils.Pending.String()
	order.Price = product.Price
	order.TotalPrice = float64(float32(i.Quantity) * product.Price)
	order.BuyerId = uid
	order.SellerId = product.SellerID
	return order
}

func (service *OrderServiceImpl) GetOrderById(id uint) model.Order {
	return service.OrderRepository.FindById(id)
}

func (service *OrderServiceImpl) GetOrderListByBuyerId(buyerId uint, limit int, offset int) []model.Order {
	return service.OrderRepository.FindByBuyerId(buyerId, limit, offset)
}

func (service *OrderServiceImpl) GetOrderListBySellerId(sellerId uint, limit int, offset int) []model.Order {
	return service.OrderRepository.FindBySellerId(sellerId, limit, offset)
}

func (service *OrderServiceImpl) UpdateStatusToAccepted(order model.Order) (model.Order, error) {
	return service.OrderRepository.UpdateStatusToAccepted(order)
}
