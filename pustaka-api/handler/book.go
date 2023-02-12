package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{
		bookService: bookService,
	}
}


func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"datas": booksResponse,
	})
}
func (h *bookHandler) GetBook(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	b, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	bookResponse := convertToResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (h *bookHandler) CreateBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
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

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToResponse(book),
	})

}

func (h *bookHandler) UpdateBooksHandler(c *gin.Context) {
	var bookRequestUpdate book.BookRequestUpdate

	err := c.ShouldBindJSON(&bookRequestUpdate)
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

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	book, err := h.bookService.Update(id, bookRequestUpdate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToResponse(book),
	})

}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	b, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToResponse(b),
	})

}

func convertToResponse(b book.Book) book.BookResponse {
	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
	}
	return bookResponse
}
