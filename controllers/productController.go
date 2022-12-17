package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/highergear/go-ecommerence/models"
	"github.com/highergear/go-ecommerence/utils"
	"net/http"
)

type CreateProductInput struct {
	Name        string  `json:"product_name" binding:"required"`
	Description string  `json:"product_description"`
	Price       float32 `json:"price" binding:"required"`
}

func CreateProduct(c *gin.Context) {
	uid, role, err := utils.ExtractTokenID(c)
	if err != nil || role != utils.Seller.String() {
		errString := "Buyer account is unauthorized to create/add products"
		if err != nil {
			errString = err.Error()
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}

	var input CreateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{}
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.SellerID = uid

	savedProduct, err := product.SaveProduct()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedProduct)
}

func GetProducts(c *gin.Context) {
	limit, offset := utils.GetLimitAndOffset(c)
	productList := models.GetProducts(limit, offset)
	c.JSON(http.StatusOK, productList)
}

func GetProductsBySellerId(c *gin.Context) {
	uid, role, err := utils.ExtractTokenID(c)
	if err != nil || role != utils.Seller.String() {
		errString := "Buyer account is unauthorized to get seller's products"
		if err != nil {
			errString = err.Error()
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}

	limit, offset := utils.GetLimitAndOffset(c)

	productList := models.GetProductsBySellerId(uid, limit, offset)

	c.JSON(http.StatusOK, productList)
}
