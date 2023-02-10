package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	dbURL := "postgres://postgres:password@localhost:5432/db_barang"
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// database.AutoMigrate(&Barang{})

	return database

}
