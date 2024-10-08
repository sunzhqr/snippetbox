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

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	fmt.Println(fileServer)

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Starting server on http://localhost:4000/")
	err := http.ListenAndServe(":4000", mux) // Web Server
	log.Fatal(err)
}
