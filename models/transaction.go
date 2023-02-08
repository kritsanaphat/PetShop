package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionID     uuid.UUID `json:"transactionID" gorm:"primary_key;"`
	DriverID          string    `json:"driverID"`
	AccountID         string    `json:"accountID"`
	ItemID            string    `json:"itemID"`
	Price             float32   `json:"price"`
	Payment           string    `json:"payment"`
	CouponID          string    `json:"couponID"`
	PaymentStatus     string    `json:"paymentStatus"`
	DeliveryStatus    string    `json:"deliveryStatus"`
	TransactionStatus string    `json:"purchaseStatus"`
}
