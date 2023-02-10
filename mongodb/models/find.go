package models

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Find() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(Ctx, bson.M{})
	if err != nil {
		
	}


}
