package repository

import "github.com/highergear/go-ecommerence/model"

type SellerRepository interface {
	Save(buyer model.Seller) (model.Seller, error)
	FindByEmail(email string) (model.Seller, error)
}
