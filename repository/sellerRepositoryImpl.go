package repository

import "github.com/highergear/go-ecommerence/model"

type SellerRepositoryImpl struct {
}

func NewSellerRepository() SellerRepository {
	return &SellerRepositoryImpl{}
}

func (s SellerRepositoryImpl) Save(seller model.Seller) (model.Seller, error) {
	var err error
	err = DB.Create(&seller).Error
	if err != nil {
		return model.Seller{}, err
	}
	return seller, nil
}

func (s SellerRepositoryImpl) FindByEmail(email string) (model.Seller, error) {
	var err error
	seller := model.Seller{}
	err = DB.Model(model.Seller{}).Where("email = ?", email).Take(&seller).Error
	if err != nil {
		return model.Seller{}, err
	}
	return seller, nil
}
