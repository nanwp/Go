package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/entity"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Name":   "Nanda",
		"Status": "Single",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hello": "Status OK",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func PostBooksHandler(c *gin.Context) {
	var bookInput entity.BookInput

	err := c.ShouldBindJSON(&bookInput)
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

	c.JSON(http.StatusOK, gin.H{
		"email": bookInput.Email,
		"title": bookInput.Title,
		"Price": bookInput.Price,
	})

}
