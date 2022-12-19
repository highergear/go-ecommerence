package repository

import (
	"github.com/highergear/go-ecommerence/model"
	"github.com/highergear/go-ecommerence/utils"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository OrderRepositoryImpl) Save(order model.Order) (model.Order, error) {
	err := DB.Create(&order).Error
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (repository OrderRepositoryImpl) FindById(id uint) model.Order {
	var order model.Order
	err := DB.Model(model.Order{}).Where("id = ?", id).Take(&order).Error
	if err != nil {
		return model.Order{}
	}

	return order
}

func (repository OrderRepositoryImpl) FindBySellerId(sellerId uint, limit int, offset int) []model.Order {
	var orderList []model.Order
	err := DB.Model(model.Order{}).Where("seller_id = ?", sellerId).Limit(limit).Offset(offset).Find(&orderList).Error
	if err != nil {
		return []model.Order{}
	}

	return orderList
}

func (repository OrderRepositoryImpl) FindByBuyerId(buyerId uint, limit int, offset int) []model.Order {
	var orderList []model.Order
	err := DB.Model(model.Order{}).Where("buyer_id = ?", buyerId).Limit(limit).Offset(offset).Find(&orderList).Error
	if err != nil {
		return []model.Order{}
	}

	return orderList
}

func (repository OrderRepositoryImpl) UpdateStatusToAccepted(order model.Order) (model.Order, error) {
	order.Status = utils.Accepted.String()
	err := DB.Save(&order).Error
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}
