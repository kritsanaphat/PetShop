package db

import (
	"log"
	"os"

	"github.com/kritsanaphat/PetShop/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	url := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&models.Account{},
		&models.Pet{},
		&models.Address{},
		&models.Shop{},
		&models.Theme{},
		&models.Comment{},
		&models.Chart{},
	)

	return db
}
