package main

import (
	"log"
	"net/http"
	"website/handle"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handle.HomeHandle)
	mux.HandleFunc("/addmahasiswa", handle.AddMahasiswa)
	mux.HandleFunc("/process", handle.Process)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Start web on port 8080")

	err := http.ListenAndServe(
		"localhost:8080",
		mux,
	)

	log.Fatal(err)
}
