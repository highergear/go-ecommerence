package repository

import (
	"github.com/highergear/go-ecommerence/model"
)

type BuyerRepository interface {
	Save(buyer model.Buyer) (model.Buyer, error)
	FindByEmail(email string) (model.Buyer, error)
}
