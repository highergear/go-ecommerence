package service

import (
	"github.com/highergear/go-ecommerence/model"
	"github.com/highergear/go-ecommerence/model/input"
	"github.com/highergear/go-ecommerence/repository"
	"github.com/highergear/go-ecommerence/utils"
	"golang.org/x/crypto/bcrypt"
)

type SellerServiceImpl struct {
	SellerRepository repository.SellerRepository
}

func NewSellerService(SellerRepository repository.SellerRepository) SellerService {
	return &SellerServiceImpl{
		SellerRepository: SellerRepository,
	}
}

func (service SellerServiceImpl) Create(i input.RegisterSellerInput) (model.Seller, error) {
	seller := model.Seller{}
	seller.Name = i.Name
	seller.Email = i.Email
	seller.Password = i.Password
	seller.AlamatPickUp = i.AlamatPickup

	savedSeller, err := service.SellerRepository.Save(seller)
	if err != nil {
		return model.Seller{}, err
	}
	return savedSeller, nil
}

func (service SellerServiceImpl) SellerLogin(email string, password string) (string, error) {
	seller, err := service.SellerRepository.FindByEmail(email)

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
