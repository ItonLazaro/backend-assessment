package config

import (
	"example/todo-go/src/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// establish MySQL connection from what was declared in .env
func ConnectDB() *gorm.DB {
	errorENV := godotenv.Load()

	if errorENV != nil {
		panic("Failed to load .env file. Kindly recheck.")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	//build a mysql request from .env values
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, errorDB := gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if errorDB != nil {
		panic("Cannot connect to database using given parameter. Kindly recheck")
	}

	Users := models.Users{}
	Tasks := models.Tasks{}

	db.AutoMigrate(&Users, &Tasks)

	return db
}

func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to disconnect to database.")
	}

	dbSQL.Close()
}
