package main

import (
	"fmt"
	"html/template"
	"net/http"

	"Groupie" // import of the package "groupie" in order to use the functions inside api.go
)

// tmpl is an HTML template for displaying data.
var tmpl = template.Must(template.ParseFiles("template/accueil.html"))

// liste stores the initial list of groups retrieved once at startup.
var liste, err = Groupie.Begin("https://groupietrackers.herokuapp.com/api")

// main is the main function of the program.
func main() {
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("server successfully up, go to http://localhost:8080")

	// Serve static files for resources like CSS, JavaScript, etc.
	fs := http.FileServer(http.Dir("template/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// Set URL handlers for different routes.
	http.HandleFunc("/", Send)
	http.HandleFunc("/Index", Index)
	http.HandleFunc("/Accueil", Accueil)
	
	http.HandleFunc("/switch", Switch)
	http.HandleFunc("/search", Search)

	// Start the HTTP server on port 8080.
	http.ListenAndServe(":8080", nil)
}

// Send fetches the group list again and passes it to the template.
func Send(w http.ResponseWriter, r *http.Request) {
	liste, err = Groupie.Begin("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, liste)
}

// Index displays the list of groups on the index page.
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, liste)
}

// Accueil displays the list of groups on the home page.
func Accueil(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("template/accueil.html"))
	tmpl.Execute(w, liste)
}

// Switch displays details of a specific artist.
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

// Search filters groups by name, ID or member name.
func Search(w http.ResponseWriter, r *http.Request) {
	letter := string(r.FormValue("q"))
	liste = Groupie.GetGroupsByName(liste, letter)
	tmpl.Execute(w, liste)
}
