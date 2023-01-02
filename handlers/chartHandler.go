package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
)

func (h handler) GetAllItemFromChart(c *gin.Context) {
	var chart []models.Chart
	accountID := c.MustGet("AccountID").(string)
	fmt.Print("FROM GetAllItemFromChart ", accountID)
	if err := h.DB.Where("id = ?", accountID).Find(&chart).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, &chart)
}

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

func (h handler) RemoveItem(c *gin.Context) {
	var json models.Chart
	c.ShouldBindJSON(&json)

	var chart models.Chart
	h.DB.Delete(&chart, json.ChartID)

	var charts []models.Chart
	c.JSON(http.StatusCreated, &charts)

}
