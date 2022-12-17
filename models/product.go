package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"size:100; not_null" json:"product_name"`
	Description string  `gorm:"size:255;" json:"description"`
	Price       float32 `gorm:"type:decimal(20,3)" json:"price"`
	SellerID    uint    `gorm:"not_null"`
}

func (p *Product) SaveProduct() (*Product, error) {
	var err error
	err = DB.Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func GetProducts(limit int, offset int) []Product {
	var productList []Product

	err := DB.Model(Product{}).Limit(limit).Offset(offset).Find(&productList).Error

	if err != nil {
		return []Product{}
	}

	return productList
}

func GetProductsBySellerId(sellerId uint, limit int, offset int) []Product {
	var productList []Product

	err := DB.Model(Product{}).Where("seller_id = ?", sellerId).Limit(limit).Offset(offset).Find(&productList).Error

	if err != nil {
		return []Product{}
	}

	return productList
}

func GetProductById(productId uint) Product {
	var p Product

	err := DB.Model(Product{}).Where("id = ?", productId).Take(&p).Error

	if err != nil {
		return Product{}
	}

	return p
}
