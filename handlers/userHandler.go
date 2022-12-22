package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kritsanaphat/PetShop/models"
)

func (h handler) GetAllUser(c *gin.Context) {

	hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	header := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(header, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { //valid
		//fmt.Println(claims["ID"])
		var users []models.Account

		if result := h.DB.Find(&users); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &users)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "messege": err.Error()})
		return
	}

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
