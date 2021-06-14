package utils

import (
	"log"
	"os"

	"gorm.io/gorm"

	// import driver
	"gorm.io/driver/mysql"
)

// Connect to database and return
func DBConnect() *gorm.DB {
	dbname := os.Getenv("DATABASE_NAME")
	uname := os.Getenv("DATABASE_USERNAME")
	upass := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")

	if host == "" {
		host = "localhost"
	}

	dsn := uname + ":" + upass + "@(" + host + ")/" + dbname + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db.Debug()
}
