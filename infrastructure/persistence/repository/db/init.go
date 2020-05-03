package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql" //for import mysql
	"github.com/joho/godotenv"
)

var dB *gorm.DB

// DBInit create connection to database
func DBInit() *gorm.DB {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", username, password, host, port, dbName, "Asia%2FJakarta")

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Panic("failed to connect to database")
	}
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(5)

	db.LogMode(true)

	dB = db
	return dB
}

// GetDB getdb
func GetDB() *gorm.DB {
	return dB
}
