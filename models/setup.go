package models

import (
	"fmt"
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

	DB.AutoMigrate(&Buyer{})
	DB.AutoMigrate(&Seller{})
	DB.AutoMigrate(&Product{})
	DB.AutoMigrate(&Order{})
}
