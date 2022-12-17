package models

import (
	"github.com/highergear/go-ecommerence/utils"
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

func (s *Seller) SaveSeller() (*Seller, error) {
	var err error
	err = DB.Create(&s).Error
	if err != nil {
		return &Seller{}, err
	}
	return s, nil
}

func (s *Seller) BeforeSave() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	s.Password = string(hashedPass)
	return nil
}

func SellerLogin(email string, password string) (string, error) {

	seller := Seller{}

	err := DB.Model(Seller{}).Where("email = ?", email).Take(&seller).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(seller.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateJwtToken(seller.ID, seller.Email, utils.Seller.String())
	if err != nil {
		return "", err
	}
	return token, nil
}
