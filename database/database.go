package database

import (
	"fmt"
	"log"
	"os"

	"github.com/itoldthekettleoff/go-forum/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	address := os.Getenv("DBADDRESS")
	name := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, address, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the DB.")
	}
	log.Println("Connected successfully")

	db.AutoMigrate(new(model.Blog))
	DBConn = db
}
