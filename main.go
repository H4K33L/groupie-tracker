package main

import (
	"fmt"
	"groupietracker/handler"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/artist", handler.ArtistHandler)
	http.HandleFunc("/artistsearch", handler.ResearchHandler)
	http.HandleFunc("/facebook", handler.FacebookHandler)
	http.HandleFunc("/twitter", handler.TwitterHandler)
	http.HandleFunc("/snapchat", handler.SnapHandler)
	http.HandleFunc("/instagram", handler.InstaHandler)
	
	fmt.Println("server successfully up, go to http://127.0.0.1:5500")
	http.ListenAndServe("127.0.0.1:5500", nil)
}
