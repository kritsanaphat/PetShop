package db

import (
	"log"

	"github.com/kritsanaphat/PetShop/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	url := "host=localhost user=postgres password=kritsanaphat dbname=petShop port=5432"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Pet{},
		&models.Address{},
	)

	return db
}
