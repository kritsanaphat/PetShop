package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) GetAllUser(c *gin.Context) {
	var users []models.User

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
	var user models.User
	if err := h.DB.Where("ID = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, &user)

}

func (h handler) Register(c *gin.Context) {
	var json models.Register
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
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := models.User{
		ID:       uuid,
		Password: string(encryptedPassword),
		Fullname: json.Fullname,
		Email:    json.Email,
	}
	address := models.Address{
		AddressID: uuid,
		Fullname:  json.Fullname,
	}
	if result := h.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	if result := h.DB.Create(&address); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &user)
	c.JSON(http.StatusCreated, &address)

}

func (h handler) Login(c *gin.Context) {
	var json models.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	var userExist models.User
	if err := h.DB.Where("Email = ?", json.Email).First(&userExist).Error; err != nil {
		fmt.Print(userExist)
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Does Not Exist"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Login Fail"})
		return

	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Login Success"})
	}

}
