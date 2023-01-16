package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
)

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
	shopID := c.MustGet("ShopID").(string)
	fmt.Print("FROM add pet shopID", shopID)
	// I can't use FromString  Method from  "github.com/gofrs/uuid"
	//uuidAccount := uuid.FromString(accountID)
	pet := models.Pet{
		ShopID:      shopID,
		ID:          uuid,
		Type:        json.Type,
		Species:     json.Species,
		Color:       json.Color,
		Sex:         json.Sex,
		Weight:      json.Weight,
		Height:      json.Height,
		Price:       json.Price,
		Description: json.Description,
		Age:         json.Age,
		Tag:         json.Tag,
	}

	if result := h.DB.Create(&pet); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &pet)
}

func (h handler) EditPet(c *gin.Context) {

}
