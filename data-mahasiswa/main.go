package main

import (
	"data-mahasiswa/handle"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", handle.HomeHandle)
	mux.HandleFunc("/detailmahasiswa", handle.DetailHandle)

	fileServer := http.FileServer(http.Dir("assets"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Start 8080")

	err := http.ListenAndServe(
		":8080",
		mux,
	)

	log.Fatal(err)
}

