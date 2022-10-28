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

func (h handler) AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	user = models.User{ID: uuid}
	json.Unmarshal(body, &user)
	result := h.DB.Create(&user)
	if (result.Error) != nil {
		fmt.Print(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}
