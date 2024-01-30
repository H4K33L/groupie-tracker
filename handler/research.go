package handler

import (
	"html/template"
	"net/http"
)

var accueil = template.Must(template.ParseFiles("templates/accueil.html"))

func ResearchHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/templates/home.html", http.StatusFound)

	err := accueil.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
