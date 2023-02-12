package main

import (
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/middleware"
	"pustaka-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := models.Connect()

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	api := router.Group("/api", gin.BasicAuth(middleware.GetUser()))

	api.GET("/books/:id", bookHandler.GetBook)
	api.GET("/books", bookHandler.GetBooks)
	api.POST("/books", bookHandler.CreateBooksHandler)
	api.PUT("/books/:id", bookHandler.UpdateBooksHandler)
	api.DELETE("/books/:id", bookHandler.DeleteBook)
	router.Run()
}
