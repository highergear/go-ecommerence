package model

import (
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

func (b *Buyer) BeforeSave() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	b.Password = string(hashedPass)
	return nil
}
