package service

import (
	"github.com/highergear/go-ecommerence/model"
	"github.com/highergear/go-ecommerence/model/input"
)

type BuyerService interface {
	Create(i input.RegisterBuyerInput) (model.Buyer, error)
	BuyerLogin(email string, password string) (string, error)
}
