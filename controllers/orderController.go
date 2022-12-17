package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/highergear/go-ecommerence/models"
	"github.com/highergear/go-ecommerence/utils"
	"net/http"
)

type CreateOrderInput struct {
	DeliverySourceAddress string `json:"source_address" binding:"required"`
	DeliveryDestAddress   string `json:"destination_address" binding:"required"`
	Items                 uint   `json:"item_id" binding:"required"`
	Quantity              int    `json:"quantity" binding:"required"`
}

func CreateOrder(c *gin.Context) {
	uid, role, err := utils.ExtractTokenID(c)
	if err != nil || role != utils.Buyer.String() {
		errString := "Seller account is unauthorized to create order"
		if err != nil {
			errString = err.Error()
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}

	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := MapInputToOrder(input)
	product := models.GetProductById(input.Items)
	if product.ID == 0 {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Product with ID: %d is not found", input.Items)})
		return
	}
	order.Price = product.Price
	order.TotalPrice = float64(float32(input.Quantity) * product.Price)
	order.BuyerId = uid
	order.SellerId = product.SellerID

	savedOrder, err := order.SaveOrder()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedOrder)
}

func MapInputToOrder(input CreateOrderInput) models.Order {
	order := models.Order{}
	order.DeliverySourceAddress = input.DeliverySourceAddress
	order.DeliveryDestAddress = input.DeliveryDestAddress
	order.Items = input.Items
	order.Quantity = input.Quantity
	order.Status = utils.Pending.String()
	return order
}

func GetOngoingOrderByBuyerId(c *gin.Context) {
	uid, _, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limit, offset := utils.GetLimitAndOffset(c)
	orderList := models.GetOrderByBuyerId(uid, limit, offset)
	c.JSON(http.StatusOK, orderList)
}

func GetOrderBySellerId(c *gin.Context) {
	uid, _, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limit, offset := utils.GetLimitAndOffset(c)
	orderList := models.GetOrderBySellerId(uid, limit, offset)
	c.JSON(http.StatusOK, orderList)
}
