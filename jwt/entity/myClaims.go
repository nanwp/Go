package entity

import "github.com/golang-jwt/jwt/v4"

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Group    string `json:"Group"`
}
