package main

import (
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	tbl := models.Connect()

	bookRepository := book.NewRepository(tbl)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	api.GET("/books/:id", bookHandler.GetBook)
	api.GET("/books", bookHandler.GetBooks)
	api.POST("/books", bookHandler.CreateBooksHandler)
	api.PUT("/books/:id", bookHandler.UpdateBooksHandler)
	api.DELETE("/books/:id", bookHandler.DeleteBook)
	router.Run()
}
