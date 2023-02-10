package entity

import "encoding/json"

type BookInput struct {
	Email string      `json:"email" binding:"required,email"`
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}
