package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kritsanaphat/PetShop/db"
	"github.com/kritsanaphat/PetShop/handlers"
	"github.com/kritsanaphat/PetShop/middleware"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := db.Init()
	h := handlers.New(DB)

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/getUserByID/:ID", h.GetProfile)
	r.POST("/register", h.UserRegister)
	r.POST("/login", h.UserLogin)
	r.POST("/adminRegister", h.AdminRegister)

	userAuthorized := r.Group("/user", middleware.UserMiddlewareJWT())
	userAuthorized.GET("getAllUser", h.GetAllUser)
	userAuthorized.GET("/profile", h.GetProfile)
	userAuthorized.GET("/theme", h.GetTheme)
	userAuthorized.POST("/shopRegister", h.ShopRegister)
	userAuthorized.GET("/swapToShop", h.LoginShop)
	userAuthorized.PATCH("/updateAddress", h.UpdateAddress)
	userAuthorized.POST("/createTheme", h.CreateTheme)
	userAuthorized.POST("/createComment", h.CreateComment)
	userAuthorized.POST("/addToChart", h.AddToChart)
	userAuthorized.DELETE("/removeItem", h.RemoveItem)
	userAuthorized.GET("/getAllItemFromChart", h.GetAllItemFromChart)

	shopAuthorized := r.Group("/shop", middleware.ShopMiddlewareJWT())
	shopAuthorized.POST("/addpet", h.AddPet)

	http.ListenAndServe(":8080", r)
}
