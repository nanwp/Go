package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/users"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUser() gin.Accounts {
	user := make(gin.Accounts)

	file, _ := h.userService.FindAll()

	for _, s := range file {
		user[s.Username] = s.Password
		// pwDecrypt := middleware.HashAndSalt(pwEncrypt)
	}
	return user
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var userRequest users.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {

		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	user, err := h.userService.Create(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
