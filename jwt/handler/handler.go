package handler

import (
	"encoding/json"
	"fmt"
	"jwt/app"
	"jwt/auth"
	"jwt/entity"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Error http method", http.StatusBadRequest)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Invalid username or Password", http.StatusBadRequest)
		return
	}

	ok, userInfo := auth.AuthenticateUser(username, password)
	if !ok {
		http.Error(w, "USER / PW SALAH", http.StatusBadRequest)
		return
	}

	// type MyClaims struct {
	// 	jwt.StandardClaims
	// 	Username string `json:"Username"`
	// 	Email    string `json:"Email"`
	// 	Group    string `json:"Group"`
	// }

	claims := entity.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    app.APPLICATION_NAME,
			ExpiresAt: time.Now().Add(app.LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: userInfo["username"].(string),
		Email:    userInfo["email"].(string),
		Group:    userInfo["group"].(string),
	}

	token := jwt.NewWithClaims(
		app.JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(app.JWT_SIGNATURE_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, _ := json.Marshal(app.M{"token": signedToken})
	w.Write([]byte(tokenString))

}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	message := fmt.Sprintf("hello %s (%s)", userInfo["Username"], userInfo["Group"])
	w.Write([]byte(message))
}
