package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
)

func (h handler) AddCoupon(c *gin.Context) {
	var json models.Coupon

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalln()
	}

	adminID := c.MustGet("AdminID")
	fmt.Print("FROM add coupon AdminID", adminID)

	coupon := models.Coupon{
		CouponID:   uuid,
		CouponCode: json.CouponCode,
		CouponName: json.CouponName,
		ExpireTime: json.ExpireTime,
		Discount:   json.Discount,
		MinPrice:   json.MinPrice,
		Condition:  json.Condition,
	}

	if result := h.DB.Create(&coupon); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &coupon)
}
