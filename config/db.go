package config

import (
	"final_project/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	hostDB = "localhost"
	userDB = "postgres"
	passDB = "kresnavw210101"
	portDB = 5432
	nameDB = "AuthAPI"
)

func StartDB() *gorm.DB {
	connectDB := fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable", hostDB, portDB, userDB, passDB, nameDB)
	db, err := gorm.Open(postgres.Open(connectDB), &gorm.Config{})

	db.Debug().AutoMigrate(&entity.User{}, &entity.Photo{})

	if err != nil {
		log.Fatal(err.Error())
	}
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	defer fmt.Println("Successfully Connected to Database")
	return db
}
