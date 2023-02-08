package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
)

func (h handler) BuyItem(c *gin.Context) {
	var json models.Transaction

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalln()
	}

	accountID := c.MustGet("AccountID").(string)
	fmt.Print("FROM buy AccountID", accountID)

	transaction := models.Transaction{
		TransactionID:     uuid,
		AccountID:         accountID,
		DriverID:          json.DriverID,
		ItemID:            json.ItemID,
		Price:             json.Price,
		Payment:           json.Payment,
		CouponID:          json.CouponID,
		PaymentStatus:     json.PaymentStatus,
		DeliveryStatus:    json.DeliveryStatus,
		TransactionStatus: json.TransactionStatus,
	}

	if result := h.DB.Create(&transaction); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &transaction)
}
