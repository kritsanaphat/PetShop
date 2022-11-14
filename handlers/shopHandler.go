package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
)

func (h handler) GetAllPet(c *gin.Context) {
	var pets []models.Pet

	if result := h.DB.Find(&pets); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &pets)
}

func (h handler) AddPet(c *gin.Context) {
	var json models.Pet
	if err := c.ShouldBindJSON(&json); err != nil { //Check the integrity of the information
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}

	pet := models.Pet{
		ID:      uuid,
		PetType: json.PetType,
		Color:   json.Color,
		Species: json.Species,
		Price:   json.Price,
		Detail:  json.Detail,
	}

	if result := h.DB.Create(&pet); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &pet)

}
