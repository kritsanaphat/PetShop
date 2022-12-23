package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritsanaphat/PetShop/models"
)

func (h handler) GetAllUser(c *gin.Context) {

	var users []models.Account

	if result := h.DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &users)

}

func (h handler) GetUserByID(c *gin.Context) {
	id := c.Params.ByName("ID")
	var user models.Account
	if err := h.DB.Where("ID = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, &user)

}

func (h handler) UpdateAddress(c *gin.Context) {
	var json models.Address
	id := c.Params.ByName("ID")
	var address models.Address
	if err := h.DB.Where("Address_ID = ?", id).First(&address).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	} else {
		fmt.Println("Found")
	}
	h.DB.Model(address).Where("Address_ID = ?", id).Updates(models.Address{
		ID:          address.ID,
		House:       json.House,
		District:    json.District,
		Subdistrict: json.Subdistrict,
		City:        json.City,
		Postcode:    json.Postcode,
	})

	c.BindJSON(&address)
	h.DB.Save(&address)
	c.JSON(http.StatusCreated, &address)
}
