package app

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type M map[string]interface{}

var APPLICATION_NAME = "Latihan JWT"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret")
