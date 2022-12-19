package repository

import (
	"github.com/highergear/go-ecommerence/model"
)

type BuyerRepositoryImpl struct {
}

func NewBuyerRepository() BuyerRepository {
	return &BuyerRepositoryImpl{}
}

func (repository *BuyerRepositoryImpl) Save(buyer model.Buyer) (model.Buyer, error) {
	var err error
	err = DB.Create(&buyer).Error
	if err != nil {
		return model.Buyer{}, err
	}
	return buyer, nil
}

func (repository *BuyerRepositoryImpl) FindByEmail(email string) (model.Buyer, error) {
	var err error
	buyer := model.Buyer{}
	err = DB.Model(model.Buyer{}).Where("email = ?", email).Take(&buyer).Error
	if err != nil {
		return model.Buyer{}, err
	}
	return buyer, nil
}
