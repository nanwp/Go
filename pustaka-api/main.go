package main

import (
	"encoding/base64"
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/middleware"
	"pustaka-api/models"
	"pustaka-api/users"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := models.Connect()

	userRepo := users.NewRepository(db)
	userService := users.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	api := router.Group("/api", basiAuth())

	api.POST("/users", userHandler.CreateUserHandler)
	api.GET("/books/:id", bookHandler.GetBook)
	api.GET("/books", bookHandler.GetBooks)
	api.POST("/books", bookHandler.CreateBooksHandler)
	api.PUT("/books/:id", bookHandler.UpdateBooksHandler)
	api.DELETE("/books/:id", bookHandler.DeleteBook)
	// router.GET("/test", basiAuth())
	router.Run()
}
func basiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		userRepo := users.NewRepository(models.Connect())
		userService := users.NewService(userRepo)
		userHandler := handler.NewUserHandler(userService)

		
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		// fmt.Printf("auth: %v\n", auth)
		if len(auth)!= 2 || auth[0] != "Basic"{
			respondWithError(401, "Unauthorized", c)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		// fmt.Println(payload)
		pair := strings.SplitN(string(payload), ":", 2)
		// fmt.Println(pair)

		a := userHandler.GetUser()
		// fmt.Println(a)
		pw := middleware.GetPwd(pair[1])


		if len(pair) != 2 || !middleware.ComparePassword(a[pair[0]], pw) {
			respondWithError(401, "Unauthorized", c)
			return
		}
		// fmt.Println("false")

		c.Next()

	}
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
