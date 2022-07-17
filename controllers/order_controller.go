package controllers

import (
	api_intergration "github.com/VuiPhan/web-service-gin/api-intergration"
	"github.com/VuiPhan/web-service-gin/db_client"
	"github.com/VuiPhan/web-service-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var order []models.Order

func CreateOrder(c *gin.Context) {
	var requestBody models.Order
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Body data is not valid",
		})
		return
	}
	requestBody.TotalAmount = caculateTotalAmount(requestBody.OrderDetail)
	res, err := db_client.DBClient.Exec("INSERT INTO OrderHeader(Notes,CreatedBy,CreatedOn) VALUES (?,?,Now());",
		requestBody.Notes,
		"App",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
		})
	}
	id, _ := res.LastInsertId()
	requestBody.OrderID = int(id)
	createOrderDetail(requestBody.OrderDetail, id)
	insertInfoCustomer(requestBody.InfoCustomer, id)
	insertOrderPayment(requestBody)
	resultAPIZalo := api_intergration.CallApiZaloPayToCreateOrder(requestBody)
	c.JSON(http.StatusCreated, models.ResponseCreateOrder{
		OrderID:              requestBody.OrderID,
		InfoOrderFromZaloPay: resultAPIZalo,
	})
}
func createOrderDetail(orderDetail []models.OrderDetail, orderID int64) bool {
	for _, element := range orderDetail {
		db_client.DBClient.Exec("INSERT INTO OrderDetail("+
			"OrderID,ProductID,"+
			"ProductName,"+
			"Price,"+
			"Amount,Notes,"+
			"CreatedBy"+
			",CreatedOn) VALUES (?,?,?,?,?,?,?,Now());",
			orderID,
			element.ProductID,
			element.ProductName,
			element.Price,
			element.Amount,
			element.Notes,
			"App",
		)
	}
	return true
}
func insertInfoCustomer(infoCustomer models.InfoCustomer, orderID int64) bool {
	_, err := db_client.DBClient.Exec("INSERT INTO "+
		"InfoCustomer(OrderID,CustomerName,Address,PhoneNumber,CreatedBy,CreatedOn) "+
		"VALUES (?,?,?,?,?,Now());",
		orderID,
		infoCustomer.CustomerName,
		infoCustomer.Address,
		infoCustomer.PhoneNumber,
		"App",
	)
	if err != nil {

	} else {

	}
	return true
}
func caculateTotalAmount(orderDetail []models.OrderDetail) float64 {
	var totalAmount float32 = 0
	for i := 0; i < len(orderDetail); i++ {
		amount := orderDetail[i].Amount * orderDetail[i].Price
		totalAmount += amount
	}
	return float64(totalAmount)
}
func insertOrderPayment(order models.Order) bool {
	_, err := db_client.DBClient.Exec("INSERT INTO "+
		"OrderPayment(OrderID,PaymentAmount,CreatedBy,CreatedOn,Status) "+
		"VALUES (?,?,?,Now(),?);",
		order.OrderID,
		order.TotalAmount,
		"App",
		0,
	)
	if err != nil {

	} else {

	}
	return true
}

func GetInfoPaymentOfOrder(c *gin.Context) {
	paramsQuery := c.Request.URL.Query()
	orderID := paramsQuery["orderID"]
	if orderID != nil {
		var orderPayment models.OrderPayment
		db_client.DBClient.Get(&orderPayment, "SELECT OrderID,PaymentAmount,CreatedBy,CreatedOn,Status FROM OrderPayment WHERE OrderID = ?", orderID[0])
		c.JSON(http.StatusOK, orderPayment)
		return
	} else {
	}
}
