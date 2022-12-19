package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Seller struct {
	gorm.Model
	Email        string    `gorm:"size:45; not_null; unique" json:"email"`
	Name         string    `gorm:"size:45; not_null;" json:"name"`
	Password     string    `gorm:"size:255; not_null;" json:"password"`
	AlamatPickUp string    `gorm:"size:255;" json:"alamat_pickup"`
	Products     []Product `gorm:"foreignKey:SellerID"`
}

func (s *Seller) BeforeSave() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	s.Password = string(hashedPass)
	return nil
}
