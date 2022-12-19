package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/highergear/go-ecommerence/controller"
	"github.com/highergear/go-ecommerence/middlewares"
	"github.com/highergear/go-ecommerence/repository"
	"github.com/highergear/go-ecommerence/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Can not load config file:", err)
	}

	repository.ConnectToDb(config)

	router := gin.Default()

	public := router.Group("/t1")
	public.POST("/buyer/register", controller.RegisterBuyer)
	public.POST("/buyer/login", controller.BuyerLogin)
	public.POST("/seller/register", controller.RegisterSeller)
	public.POST("/seller/login", controller.SellerLogin)
	public.GET("product", controller.GetProductList)

	protected := router.Group("/t2")
	protected.Use(middlewares.JwtAuthenticateMiddleware())
	protected.POST("/product", controller.CreateProduct)
	protected.GET("/seller/product", controller.GetProductsBySellerId)
	protected.GET("/seller/order", controller.GetOrderBySellerId)
	protected.PUT("/seller/order", controller.UpdateOrderStatusToAccepted)
	protected.GET("/buyer/order", controller.GetOngoingOrderByBuyerId)
	protected.POST("/buyer/order", controller.CreateOrder)

	err = router.Run("localhost:8080")
	if err != nil {
		return
	}
}
