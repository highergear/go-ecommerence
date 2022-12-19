package model

import (
	"github.com/jinzhu/gorm"
)

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
