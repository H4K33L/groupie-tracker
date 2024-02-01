package main

import (
	"fmt"
    "html/template"
    "net/http"

	"Groupie"
)

var tmpl = template.Must(template.ParseFiles("Data/Page-Hangman.html"))

var liste, err = Groupie.Begin("https://groupietrackers.herokuapp.com/api")

var toSend = liste

func main() {
	if err != nil {
		fmt.Println(err)
		return
	}

    fmt.Println("server succefully up, go to http://localhost:8080")
    http.HandleFunc("/",Send)
    http.HandleFunc("/switch",Switch)
    http.HandleFunc("/switchbis",SwitchBis)
    http.ListenAndServe(":8080", nil)
}

func Send(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, toSend)
}

func Switch(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("Data/pagebis.html"))
    letter := string(r.FormValue("name"))
	fmt.Println(letter)
	tmpl.Execute(w, toSend)
}

func SwitchBis(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("Data/Page-Hangman.html"))
	tmpl.Execute(w, toSend)
}