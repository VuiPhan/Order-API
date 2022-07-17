package models

import "time"

type Order struct {
	OrderID      int           `json:"order_id"`
	Notes        string        `json:"notes"`
	TotalAmount  float64       `json:"total_amount"`
	OrderDetail  []OrderDetail `json:"order_detail"`
	CreatedBy    string        `json:"created_by"`
	CreatedOn    time.Time     `json:"created_on"`
	InfoCustomer InfoCustomer  `json:"info_customer"`
}
type OrderDetail struct {
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float32 `json:"price"`
	Amount      float32 `json:"amount"`
	Notes       string  `json:"notes"`
}
type ResponseCreateOrder struct {
	OrderID              int                            `json:"order_id"`
	InfoOrderFromZaloPay ResponseCreateOrderFromZaloPay `json:"info_order_from_zalo_pay"`
}

type ResponseCreateOrderFromZaloPay struct {
	OrderToken       string `json:"order_token"`
	OrderUrl         string `json:"order_url"`
	ReturnCode       int    `json:"return_code"`
	ReturnMessage    string `json:"return_message"`
	SubReturnCode    int    `json:"sub_return_code"`
	SubReturnMessage string `json:"sub_return_message"`
	ZpTransToken     string `json:"zp_trans_token"`
}
