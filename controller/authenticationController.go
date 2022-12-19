package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/highergear/go-ecommerence/model/input"
	"github.com/highergear/go-ecommerence/repository"
	"github.com/highergear/go-ecommerence/service"
	"net/http"
)

var buyerRepository = repository.NewBuyerRepository()
var buyerService = service.NewBuyerService(buyerRepository)
var sellerRepository = repository.NewSellerRepository()
var sellerService = service.NewSellerService(sellerRepository)

func RegisterBuyer(c *gin.Context) {
	var i input.RegisterBuyerInput
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedBuyer, err := buyerService.Create(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedBuyer)
}

func BuyerLogin(c *gin.Context) {
	var i input.LoginInput
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := buyerService.BuyerLogin(i.Email, i.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func RegisterSeller(c *gin.Context) {
	var i input.RegisterSellerInput
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedSeller, err := sellerService.Create(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedSeller)
}

func SellerLogin(c *gin.Context) {
	var i input.LoginInput
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := sellerService.SellerLogin(i.Email, i.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
