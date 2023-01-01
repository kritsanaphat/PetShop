package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (h handler) GetProfile(c *gin.Context) {
	var user []models.Account
	accountID := c.MustGet("AccountID").(string)
	fmt.Print("FROM Profile ", accountID)
	if err := h.DB.Where("account_id = ?", accountID).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func (h handler) GetTheme(c *gin.Context) {
	var themes []models.Theme

	if result := h.DB.Find(&themes); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &themes)
}

func (h handler) UpdateAddress(c *gin.Context) {

	accountID := c.MustGet("AccountID").(string)
	var json models.Address
	c.ShouldBindJSON(&json)

	var address models.Address
	if err := h.DB.Where("id = ?", accountID).First(&address).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	} else {
		fmt.Println("Found")
	}
	h.DB.Model(address).Where("id = ?", accountID).Updates(models.Address{
		House:       json.House,
		District:    json.District,
		Subdistrict: json.Subdistrict,
		City:        json.City,
		Postcode:    json.Postcode,
	})

	h.DB.Save(&address)
	c.JSON(http.StatusCreated, &address)
}
