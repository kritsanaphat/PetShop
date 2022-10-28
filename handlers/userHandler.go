package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
)

func (h handler) GetRegister(c *gin.Context) {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &users)
}

func (h handler) Register(c *gin.Context) {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}
	var user models.User = models.User{ID: uuid}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, &user)

}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	var json models.Login
	var user models.User
	fmt.Print(json.Fullname, json.Password)
	h.DB.Where("Fullname = ?", json.Fullname).First(&user)
	h.DB.Where("Fullname = ?", json.Password).First(&user)

}
