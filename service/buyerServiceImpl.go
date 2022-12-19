package service

import (
	"github.com/highergear/go-ecommerence/model"
	"github.com/highergear/go-ecommerence/model/input"
	"github.com/highergear/go-ecommerence/repository"
	"github.com/highergear/go-ecommerence/utils"
	"golang.org/x/crypto/bcrypt"
)

type BuyerServiceImpl struct {
	BuyerRepository repository.BuyerRepository
}

func NewBuyerService(BuyerRepository repository.BuyerRepository) BuyerService {
	return &BuyerServiceImpl{
		BuyerRepository: BuyerRepository,
	}
}

func (service *BuyerServiceImpl) Create(i input.RegisterBuyerInput) (model.Buyer, error) {
	buyer := model.Buyer{}
	buyer.Name = i.Name
	buyer.Email = i.Email
	buyer.Password = i.Password
	buyer.AlamatPengiriman = i.AlamatPengiriman

	savedBuyer, err := service.BuyerRepository.Save(buyer)
	if err != nil {
		return model.Buyer{}, err
	}
	return savedBuyer, nil
}

func VerifyPassword(hashedPass string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}

func (service *BuyerServiceImpl) BuyerLogin(email string, password string) (string, error) {
	buyer, err := service.BuyerRepository.FindByEmail(email)

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
