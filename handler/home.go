package handler

import (
	"html/template"
	"net/http"
	
)

var home = template.Must(template.ParseFiles("templates/home.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	err := home.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
