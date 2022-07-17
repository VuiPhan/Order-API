package controllers

import (
	"encoding/json"
	"github.com/VuiPhan/web-service-gin/db_client"
	"github.com/VuiPhan/web-service-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/zpmep/hmacutil"
	"net/http"
	"strings"
)

var (
	key2 = "trMrHtvjo6myautxDUiAcYsVtaeQ8nhf"
)

func CallBackOrderFromZaloPay(c *gin.Context) {
	var requestBody models.InfoCallbackFromZaloPay
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Body data is not valid",
		})
		return
	}

	macCompare := hmacutil.HexStringEncode(hmacutil.SHA256, key2, requestBody.Data)
	result := make(map[string]interface{})
	if macCompare == requestBody.Mac {
		result["return_code"] = 1
		result["return_message"] = "success"
	} else {
		result["return_code"] = -1
		result["return_message"] = "mac not equal"
	}
	var infoOrderCallBack models.InfoOrderCallBack
	json.Unmarshal([]byte(requestBody.Data), &infoOrderCallBack)
	resultUpdate := updateStatusPaymentOrder(infoOrderCallBack)
	if resultUpdate == true {

	}

	c.JSON(http.StatusCreated, result)

}
func updateStatusPaymentOrder(order models.InfoOrderCallBack) bool {
	orderID := strings.Split(order.AppTransId, "_")[1]
	_, err := db_client.DBClient.Exec("UPDATE "+
		"OrderPayment "+
		"SET STATUS = 1 "+
		"WHERE OrderID = ? ",
		orderID,
	)
	if err != nil {

	} else {

	}
	return true
}
