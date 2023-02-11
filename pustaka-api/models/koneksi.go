package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	dbURL := "postgres://postgres:password@localhost:5432/go_api"
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//database.AutoMigrate(&entity.Book)

	return database

}
