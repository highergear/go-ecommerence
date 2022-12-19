package service

import (
	"github.com/highergear/go-ecommerence/model"
	"github.com/highergear/go-ecommerence/model/input"
)

type SellerService interface {
	Create(i input.RegisterSellerInput) (model.Seller, error)
	SellerLogin(email string, password string) (string, error)
}
