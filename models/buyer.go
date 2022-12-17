package models

import (
	"github.com/highergear/go-ecommerence/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Buyer struct {
	gorm.Model
	Email            string `gorm:"size:45; not_null; unique" json:"email"`
	Name             string `gorm:"size:45; not_null;" json:"name"`
	Password         string `gorm:"size:255; not_null;" json:"password"`
	AlamatPengiriman string `gorm:"size:255;" json:"alamat_pengiriman"`
}

func (b *Buyer) SaveBuyer() (*Buyer, error) {
	var err error
	err = DB.Create(&b).Error
	if err != nil {
		return &Buyer{}, err
	}
	return b, nil
}

func (b *Buyer) BeforeSave() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	b.Password = string(hashedPass)
	return nil
}

func VerifyPassword(hashedPass string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}

func BuyerLogin(email string, password string) (string, error) {

	buyer := Buyer{}

	err := DB.Model(Buyer{}).Where("email = ?", email).Take(&buyer).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(buyer.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateJwtToken(buyer.ID, buyer.Email, utils.Buyer.String())
	if err != nil {
		return "", err
	}
	return token, nil
}
