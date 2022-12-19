package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/highergear/go-ecommerence/model/input"
	"github.com/highergear/go-ecommerence/repository"
	"github.com/highergear/go-ecommerence/service"
	"github.com/highergear/go-ecommerence/utils"
	"net/http"
)

var productRepository = repository.NewProductRepository()
var productService = service.NewProductService(productRepository)

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

	var i input.CreateProductInput
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedProduct, err := productService.Create(i, uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedProduct)
}

func GetProductList(c *gin.Context) {
	limit, offset := utils.GetLimitAndOffset(c)
	productList := productService.GetProductList(limit, offset)
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
	productList := productService.GetProductListBySellerId(uid, limit, offset)

	c.JSON(http.StatusOK, productList)
}
