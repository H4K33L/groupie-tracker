package main

import (
	"fmt"
	"groupietracker/handler"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/artist", handler.ArtistHandler)
	http.HandleFunc("/contact", handler.ContactHandler)

	port := 8080

	fmt.Printf("Serveur écoutant sur le port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
