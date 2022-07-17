package controllers

import (
	"github.com/VuiPhan/web-service-gin/db_client"
	"github.com/VuiPhan/web-service-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetListProduct(c *gin.Context) {
	var listProduct []models.Product
	listProduct = append(listProduct, models.Product{ProductName: "Vui"})

	rows := db_client.DBClient.Select(&listProduct, "SELECT * FROM Products")
	if rows != nil {

	}
	c.ShouldBindJSON(listProduct)
	c.JSON(http.StatusOK, gin.H{
		"data": listProduct,
	})
}
