package main

import (
	"fmt"
	"html/template"
	"net/http"

	"Groupie"
)

var tmpl = template.Must(template.ParseFiles("template/accueil.html"))

var liste, err = Groupie.Begin("https://groupietrackers.herokuapp.com/api")

func main() {
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("server succefully up, go to http://localhost:5000")

	fs := http.FileServer(http.Dir("template/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", Send)

	http.HandleFunc("/Index", Index)
	http.HandleFunc("/Accueil", Accueil)
	http.HandleFunc("/Map", Map)
	http.HandleFunc("/switch", Switch)
	http.HandleFunc("/search", Search)

	http.ListenAndServe(":5000", nil)
}

func Send(w http.ResponseWriter, r *http.Request) {
	liste, err = Groupie.Begin("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, liste)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, liste)
}

func Accueil(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("template/accueil.html"))
	tmpl.Execute(w, liste)
}
func Map(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("template/map.html"))
	tmpl.Execute(w, liste)
}

func Switch(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("template/artistes.html"))
	letter := string(r.FormValue("name"))
	artist, err := Groupie.LoadArtist("https://groupietrackers.herokuapp.com/api", letter)
	if err != nil {
		fmt.Println(err)
		return
	}
	artist.Information, err = Groupie.GetLinkInfos(artist.Relation)
	tmpl.Execute(w, artist)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	letter := string(r.FormValue("q"))
	liste = Groupie.GetGroupsByName(liste, letter)
	tmpl.Execute(w, liste)
}
