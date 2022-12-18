package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/highergear/go-ecommerence/controllers"
	"github.com/highergear/go-ecommerence/middlewares"
	"github.com/highergear/go-ecommerence/models"
	"github.com/highergear/go-ecommerence/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Can not load config file:", err)
	}

	models.ConnectToDb(config)

	router := gin.Default()

	public := router.Group("/t1")
	public.POST("/buyer/register", controllers.RegisterBuyer)
	public.POST("/buyer/login", controllers.BuyerLogin)
	public.POST("/seller/register", controllers.RegisterSeller)
	public.POST("/seller/login", controllers.SellerLogin)
	public.GET("product", controllers.GetProducts)

	protected := router.Group("/t2")
	protected.Use(middlewares.JwtAuthenticateMiddleware())
	protected.POST("/product", controllers.CreateProduct)
	protected.GET("/seller/product", controllers.GetProductsBySellerId)
	protected.GET("/seller/order", controllers.GetOrderBySellerId)
	protected.PUT("/seller/order", controllers.UpdateOrderStatusToAccepted)
	protected.GET("/buyer/order", controllers.GetOngoingOrderByBuyerId)
	protected.POST("/buyer/order", controllers.CreateOrder)

	err = router.Run("localhost:8080")
	if err != nil {
		return
	}
}
