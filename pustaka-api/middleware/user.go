package middleware

import (
	"pustaka-api/models"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
)

func GetUser() gin.Accounts{
	users := make(gin.Accounts)

	userRepo := user.NewRepository(models.Connect())
	userService := user.NewService(userRepo)

	file, _ := userService.FindAll()

	for _, s := range file {
		users[s.Username] = s.Password
	}

	return users
}
