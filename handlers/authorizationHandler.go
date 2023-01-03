package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"

	//"github.com/google/uuid"
	"github.com/kritsanaphat/PetShop/models"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

func (h handler) UserRegister(c *gin.Context) {
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
	user := models.Account{
		AccountID: uuid,
		Password:  string(encryptedPassword),
		Username:  json.Username,
		Email:     json.Email,
		Phone:     json.Phone,
	}
	address := models.Address{
		ID: uuid,
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

func (h handler) UserLogin(c *gin.Context) {
	var json models.Login
	var userExist models.Account

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := h.DB.Where("Username = ?", json.Username).First(&userExist).Error; err != nil {
		fmt.Print(userExist)
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Does Not Exist"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Login Fail"})
		return

	} else {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"ID":  userExist.AccountID,
			"exp": time.Now().Add(time.Minute * 1).Unix(), //Exp just 1 min
		})

		tokenString, _ := token.SignedString(hmacSampleSecret)

		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Login Success", "token": tokenString})
	}
}

func (h handler) ShopRegister(c *gin.Context) {
	var json models.Shop

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
	accountID := c.MustGet("AccountID").(string)
	// I can't use FromString  Method from  "github.com/gofrs/uuid"
	//uuidAccount := uuid.FromString(accountID)

	fmt.Print("FROM ShopRegister accountID ", accountID)
	fmt.Print("FROM ShopRegister shopID ", uuid)
	shop := models.Shop{
		AccountID: accountID,
		ShopID:    uuid,
		ShopName:  json.ShopName,
		Firstname: json.Firstname,
		Lastname:  json.Lastname,
		Phone:     json.Phone,
	}

	if result := h.DB.Create(&shop); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &shop)
}

func (h handler) SwaptoShop(c *gin.Context) {

	//query user token
	var shop models.Shop
	accountID := c.MustGet("AccountID").(string)
	fmt.Println("FROM SWAP AccountID", accountID)

	if err := h.DB.Where("account_id = ?", accountID).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "The user has not registered as a shop account"})
		return
	}

	fmt.Println("SHOPID", shop.ShopID)
	hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY2"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":  shop.ShopID,
		"exp": time.Now().Add(time.Minute * 1).Unix(), //Exp just 1 min
	})

	tokenString, _ := token.SignedString(hmacSampleSecret)
	fmt.Println("FROM SWAP token shop", tokenString)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Swap to Shop Success", "token": tokenString})
}

func (h handler) AdminRegister(c *gin.Context) {
	var json models.Admin

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
	admin := models.Admin{
		AdminID:    uuid,
		EmployeeID: json.EmployeeID,
		Password:   string(encryptedPassword),
		Username:   json.Username,
		Email:      json.Email,
	}

	if result := h.DB.Create(&admin); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, "Register Admin Successfully")
	c.JSON(http.StatusCreated, &admin)

}
