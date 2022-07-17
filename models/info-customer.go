package models

type InfoCustomer struct {
	OrderID      int    `json:"order_id"`
	CustomerName string `json:"customer_name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
}
