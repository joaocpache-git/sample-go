package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sample/domain"
)

func ConnectDB() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting database %v", err)
		panic(err)
	}

	//	defer db.Close()

	db.AutoMigrate(&domain.User{})

	return db
}
