package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("heheee")
	//db_client.InitialConnection()
	router := gin.Default()
	//router.Use(middleware.CORSMiddleware())
	//router.POST("/create-order", controllers.CreateOrder)
	//router.GET("/info-order-payment", controllers.GetInfoPaymentOfOrder)
	//
	//router.GET("/get-list-product", controllers.GetListProduct)
	//router.POST("/callback", controllers.CallBackOrderFromZaloPay)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":3000")
}
