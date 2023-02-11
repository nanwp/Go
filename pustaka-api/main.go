package main

import (
	"pustaka-api/entity"
	"pustaka-api/handler"
	"pustaka-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	tbl := models.Connect().Table("tbl_book")

	bookRepository := entity.NewRepository(tbl)
	bookService := entity.NewService(bookRepository)

	bookRequest := entity.BookRequest{
		Title:       "Nanda Story Happy",
		Price:       "79000",
	}

	bookService.Create(bookRequest)

	router.GET("/", handler.RootHandler)
	router.GET("/hello", handler.HelloHandler)
	router.GET("/books/:id/:title", handler.BooksHandler)
	router.GET("/books", handler.QueryHandler)
	router.POST("/books", handler.PostBooksHandler)
	router.Run()
}
