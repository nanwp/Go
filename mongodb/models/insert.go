package models

import (
	"fmt"
	"log"
)

func Insert(s Student) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(Ctx, s)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success")
}
