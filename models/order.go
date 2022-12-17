package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	BuyerId               uint    `gorm:"not_null"`
	SellerId              uint    `gorm:"not_null"`
	DeliverySourceAddress string  `gorm:"size:255; not_null"`
	DeliveryDestAddress   string  `gorm:"size:255; not_null"`
	Items                 uint    `gorm:"not_null"`
	Quantity              int     `gorm:"not_null"`
	Price                 float32 `gorm:"not_null"`
	TotalPrice            float64 `gorm:"not_null"`
	Status                string  `gorm:"not_null"`
}

func (o *Order) SaveOrder() (*Order, error) {
	var err error
	err = DB.Create(&o).Error
	if err != nil {
		return &Order{}, err
	}
	return o, nil
}

func GetOrderBySellerId(sellerId uint, limit int, offset int) []Order {
	var orderList []Order

	err := DB.Model(Order{}).Where("seller_id = ?", sellerId).Limit(limit).Offset(offset).Find(&orderList).Error

	if err != nil {
		return []Order{}
	}

	return orderList
}

func GetOrderByBuyerId(buyerId uint, limit int, offset int) []Order {
	var orderList []Order

	err := DB.Model(Order{}).Where("buyer_id = ?", buyerId).Limit(limit).Offset(offset).Find(&orderList).Error

	if err != nil {
		return []Order{}
	}

	return orderList
}
