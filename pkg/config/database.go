package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func GetConnectionString() string {

	username := os.Getenv("DB_USER")
	if username == "" {
		username = "root"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "12345"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "go-book"
	}
	charset := os.Getenv("CHARSET")
	if charset == "" {
		charset = "utf8"
	}
	parseTime := os.Getenv("PARSE_TIME")
	if parseTime == "" {
		parseTime = "True"
	}
	loc := os.Getenv("LOCATION")
	if loc == "" {
		loc = "Local"
	}
	return fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=%s&loc=%s", username, password, dbName, charset, parseTime, loc)
}

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "mysql"
	}
	d, err := gorm.Open(driver, GetConnectionString())
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
