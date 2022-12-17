package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/highergear/go-ecommerence/models"
	"net/http"
)

type RegisterBuyerInput struct {
	Email            string `json:"email" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Password         string `json:"password" binding:"required"`
	AlamatPengiriman string `json:"alamat_pengiriman"`
}

type RegisterSellerInput struct {
	Email        string `json:"email" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Password     string `json:"password" binding:"required"`
	AlamatPickup string `json:"alamat_pickup"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterBuyer(c *gin.Context) {

	var input RegisterBuyerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buyer := models.Buyer{}
	buyer.Name = input.Name
	buyer.Email = input.Email
	buyer.Password = input.Password
	buyer.AlamatPengiriman = input.AlamatPengiriman

	savedBuyer, err := buyer.SaveBuyer()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedBuyer)
}

func BuyerLogin(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.BuyerLogin(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func RegisterSeller(c *gin.Context) {

	var input RegisterSellerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seller := models.Seller{}

	seller.Name = input.Name
	seller.Email = input.Email
	seller.Password = input.Password
	seller.AlamatPickUp = input.AlamatPickup

	savedSeller, err := seller.SaveSeller()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedSeller)
}

func SellerLogin(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.SellerLogin(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
