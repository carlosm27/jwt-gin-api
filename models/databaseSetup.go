package models

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func Setup() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	dbUrl := fmt.Sprint(os.Getenv("DATABASE_URL"))

	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	if err = db.AutoMigrate(&User{}); err != nil {
		log.Println(err)
	}

	return db, err
}
