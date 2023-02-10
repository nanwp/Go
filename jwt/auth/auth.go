package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jwt/app"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/novalagung/gubrak/v2"
)

func AuthenticateUser(username, password string) (bool, app.M) {
	basePath, _ := os.Getwd()
	dbPath := filepath.Join(basePath, "users.json")
	buf, _ := ioutil.ReadFile(dbPath)

	data := make([]app.M, 0)
	err := json.Unmarshal(buf, &data)
	if err != nil {
		return false, nil
	}

	res := gubrak.From(data).Find(func(each app.M) bool {
		return each["username"] == username && each["password"] == password

	}).Result()

	if res != nil {
		resM := res.(app.M)
		delete(resM, "password")
		return true, resM
	}

	return false, nil
}

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != app.JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return app.JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}
