package user

import (
	"encoding/base64"
	"pustaka-api/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), "", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			respondWithError(401, "Unautorized", c)
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !AutheticateUser(pair[0], pair[1]) {
			respondWithError(401, "Unauthorized", c)
			return
		}

		c.Next()
	}
}

func AutheticateUser(username, password string) bool {
	var user User

	err := models.Connect().Table("tbl_user").Where(User{Username: username, Password: password}).FirstOrCreate(&user)

	if err.Error != nil {
		return false
	}
	return true
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
