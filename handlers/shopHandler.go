package handlers

import (
	"database/sql/driver"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
)

// Value validate enum when set to database
func (t models.petType) Value() (driver.Value, error) {
	switch t {
	case IT, Decorate, Etc: //valid case
		return string(t), nil
	}
	return nil, errors.New("Invalid product type value") //else is invalid
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
		ID:     uuid,
		Price:  json.Price,
		Detail: json.Detail,
	}

	if result := h.DB.Create(&pet); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &pet)

}
