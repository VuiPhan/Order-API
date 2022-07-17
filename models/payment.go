package models

import "time"

type OrderPayment struct {
	OrderID       int       `json:"order_id" db:"OrderID"`
	PaymentAmount float32   `json:"payment_amount" db:"PaymentAmount"`
	CreatedBy     string    `json:"created_by" db:"CreatedBy"`
	CreatedOn     time.Time `json:"created_on" db:"CreatedOn"`
	Status        int       `json:"status" db:"Status"`
}
