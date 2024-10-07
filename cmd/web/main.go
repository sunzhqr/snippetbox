package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // Servemux or router
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fmt.Println("Starting server on http://localhost:4000/")
	err := http.ListenAndServe(":4000", mux) // Web Server
	log.Fatal(err)
}
