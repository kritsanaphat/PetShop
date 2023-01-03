package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	CouponID   uuid.UUID `json:"couponID" gorm:"primary_key;"`
	CouponCode string    `json:"couponCode"`
	CouponName string    `json:"couponName"`
	ExpireTime time.Time `json:"expireTime"`
	Discount   float32   `json:"discount"`
	MinPrice   float32   `json:"minPrice"`
	Condition  string    `json:"condition"`
}
