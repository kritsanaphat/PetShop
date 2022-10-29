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

func (h handler) Login(c *gin.Context) {
	var json models.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	fmt.Print(json.Email, json.Password)
	var emailExist models.Login
	if err := h.DB.Where("Email = ?", json.Email).First(&emailExist).Error; err != nil {
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Does Not Exist"})
		return
	}

	var passwordExist models.Login
	if err := h.DB.Where("Password = ?", json.Password).First(&passwordExist).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Login Success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Login Fail"})
		return
	}

}
