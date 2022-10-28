package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/models"
)

func (h handler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var user models.User = models.User{ID: uuid}
	// fmt.Println("BE", user)
	json.Unmarshal(body, &user)
	// fmt.Println("AF", user)

	result := h.DB.Create(&user) // pass pointer of data to Create
	if (result.Error) != nil {
		fmt.Print(result.Error)
	}

}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	var json models.Login
	var user models.User

	h.DB.Where("Fullname = ?", json.Fullname).First(&user)
	h.DB.Where("Fullname = ?", json.Password).First(&user)

}
