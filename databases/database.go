package databases

import (
	"log"
	"os"

	"github.com/kritsanaphat/PetShop/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// // DB is a global var for connect DB
// var DB *gorm.DB

// // DBConfig represents db configuration
// type DBConfig struct {
// 	Host     string
// 	Port     int
// 	User     string
// 	DBName   string
// 	Password string
// }

// // BuildDBConfig use for building DB config
// func BuildDBConfig() *DBConfig {
// 	dbConfig := DBConfig{
// 		Host:     "localhost",
// 		Port:     5432,
// 		User:     "postgres",
// 		DBName:   "petShop",
// 		Password: os.Getenv("DB_PASSWORD"),
// 	}
// 	return &dbConfig
// }

// // DbURL use for create DB connection URL
// func DbURL() string {
// 	return os.Getenv("DATABASE_URL")
// }

// DB is a global var for connect DB

func Init() *gorm.DB {
	url := os.Getenv("DATABASE_URL")
	print(url)
	DB, err := gorm.Open(postgres.Open("host=localhost user=postgres password=kritsanaphat dbname=petShop port=5432"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	DB.AutoMigrate(
		&models.Account{},
		&models.Pet{},
		&models.Address{},
		&models.Shop{},
		&models.Theme{},
		&models.Comment{},
		&models.Chart{},
		&models.Admin{},
		&models.Coupon{},
		&models.Transaction{},
	)

	return DB
}
