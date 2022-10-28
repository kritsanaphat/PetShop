package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kritsanaphat/PetShop/db"
	"github.com/kritsanaphat/PetShop/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()
	// p := models.User{Fullname: "krit"}
	// fmt.Print(p)
	router.HandleFunc("/user", h.AddUser).Methods(http.MethodPost)
	http.ListenAndServe(":8080", router)
}
