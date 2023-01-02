package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
)

func (h handler) AddToChart(c *gin.Context) {
	var json models.Chart
	c.ShouldBindJSON(&json)

	accountID := c.MustGet("AccountID").(string)

	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("FROM create comment addToChart", accountID)
	chart := models.Chart{
		ChartID: uuid,
		ID:      accountID,
		Item:    json.Item,
	}

	if result := h.DB.Create(&chart); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, &chart)
}
