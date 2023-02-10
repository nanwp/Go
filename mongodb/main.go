package main

import (
	"mongodb/models"
)

func main() {
	models.Insert(models.Student{Name: "nanda",Grade: 4})
}
