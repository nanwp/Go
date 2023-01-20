package main

import (
	"database/handle"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", handle.UserHandler)

	log.Println("JALAN SERVER PORT 8080")

	err := http.ListenAndServe(
		":8080",
		mux,
	)
	log.Fatal(err)

}
