package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=3306 user=root dbname=db_bolt sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("Error", err)
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxIdleTime(10)
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {
	return db
}
