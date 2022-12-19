package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"size:100; not_null" json:"product_name"`
	Description string  `gorm:"size:255;" json:"description"`
	Price       float32 `gorm:"type:decimal(20,3)" json:"price"`
	SellerID    uint    `gorm:"not_null"`
}
