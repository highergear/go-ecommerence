package repository

import (
	"fmt"
	"github.com/highergear/go-ecommerence/model"
	"github.com/highergear/go-ecommerence/utils"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDb(config utils.Config) {
	var err error
	DBUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName)

	DB, err = gorm.Open(config.DBDriver, DBUrl)

	if err != nil {
		log.Fatal("Can not connect to database:", err)
	} else {
		log.Println("Connected to database....")
	}

	DB.AutoMigrate(&model.Buyer{})
	DB.AutoMigrate(&model.Seller{})
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.Order{})
}
