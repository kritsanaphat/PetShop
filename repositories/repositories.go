package repositories

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type todoRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) todoRepository {
	return todoRepository{db}
}

func (h todoRepository) UserRegister(c *gin.Context) {
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
