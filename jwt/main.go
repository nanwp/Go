package main

import (
	"fmt"
	"jwt/auth"
	"jwt/handler"
	"jwt/middleware"
	"net/http"
)

func main() {
	mux := new(middleware.MuxServer)
	mux.RegisterMiddleware(auth.MiddlewareJWTAuthorization)

	mux.HandleFunc("/login", handler.HandlerLogin)
	mux.HandleFunc("/index", handler.HandlerIndex)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":8080"

	fmt.Println("Server at", server.Addr)
	server.ListenAndServe()

}
