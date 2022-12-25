package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
)

func (h handler) CreateTheme(c *gin.Context) {
	var json models.Theme

	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}

	accountID := c.MustGet("AccountID").(string)
	fmt.Print("FROM create theme accountID", accountID)
	theme := models.Theme{
		ThemeID:  uuid,
		AuthorID: accountID,
		Topic:    json.Topic,
		Content:  json.Content,
	}

	if result := h.DB.Create(&theme); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
}
