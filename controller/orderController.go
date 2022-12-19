package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/highergear/go-ecommerence/model/input"
	"github.com/highergear/go-ecommerence/repository"
	"github.com/highergear/go-ecommerence/service"
	"github.com/highergear/go-ecommerence/utils"
	"net/http"
	"strconv"
)

type CreateOrderInput struct {
	DeliverySourceAddress string `json:"source_address" binding:"required"`
	DeliveryDestAddress   string `json:"destination_address" binding:"required"`
	Items                 uint   `json:"item_id" binding:"required"`
	Quantity              int    `json:"quantity" binding:"required"`
}

var orderRepository = repository.NewOrderRepository()
var orderService = service.NewOrderService(orderRepository)

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

	var i input.CreateOrderInput
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := productService.GetProductById(i.Items)
	if product.ID == 0 {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Product with ID: %d is not found", i.Items)})
		return
	}

	savedOrder, err := orderService.Create(i, product, uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedOrder)
}

func GetOngoingOrderByBuyerId(c *gin.Context) {
	uid, role, err := utils.ExtractTokenID(c)
	if err != nil || role != utils.Buyer.String() {
		errString := "Seller account is unauthorized to list buyer's orders"
		if err != nil {
			errString = err.Error()
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}
	limit, offset := utils.GetLimitAndOffset(c)
	orderList := orderService.GetOrderListByBuyerId(uid, limit, offset)
	c.JSON(http.StatusOK, orderList)
}

func GetOrderBySellerId(c *gin.Context) {
	uid, role, err := utils.ExtractTokenID(c)
	if err != nil || role != utils.Seller.String() {
		errString := "Buyer account is unauthorized to list seller's orders"
		if err != nil {
			errString = err.Error()
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}
	limit, offset := utils.GetLimitAndOffset(c)
	orderList := orderService.GetOrderListBySellerId(uid, limit, offset)
	c.JSON(http.StatusOK, orderList)
}

func UpdateOrderStatusToAccepted(c *gin.Context) {
	uid, role, err := utils.ExtractTokenID(c)
	if err != nil || role != utils.Seller.String() {
		errString := "Buyer account is unauthorized to update order status to Accepted"
		if err != nil {
			errString = err.Error()
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}

	if c.Query("order_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameter: order_id"})
		return
	}

	orderId, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := orderService.GetOrderById(uint(orderId))
	if order.ID == 0 {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Order with ID: %d is not found", orderId)})
		return
	}

	if order.SellerId != uid {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Order with ID: %d is not belong to seller ID: %d", orderId, uid)})
		return
	}

	if order.Status != utils.Pending.String() {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Unable to update order with ID: %d, caused by its status is not Pending", orderId)})
		return
	}

	updatedOrder, err := orderService.UpdateStatusToAccepted(order)
	c.JSON(http.StatusOK, updatedOrder)
}
