package main

import (
	"Groupie"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type artist struct {
	ID           float64  `json:"id"`
	Img          string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate float64  `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relation     string   `json:"relation"`
}

var data []artist

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	currentDir := getCurrentDirectory()
	templatePath := currentDir + "/../template/" + tmpl + ".html"
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

func getCurrentDirectory() string {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return strings.ReplaceAll(workingDir, "\\", "/")
}

func main() {
	fs := http.FileServer(http.Dir("../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Static files served from /static/")
	start()

	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)
	fmt.Printf("Starting server to test (listen 8080)\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index")
}

func start() {
	information, err := Groupie.GetApi("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Println(err)
		return
	}

	groups, err := Groupie.GetArtist(information.Artists)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, group := range groups {
		artist := artist{
			ID:           group.ID,
			Img:          group.Img,
			Name:         group.Name,
			Members:      group.Members,
			CreationDate: group.CreationDate,
			FirstAlbum:   group.FirstAlbum,
			Locations:    group.Locations,
			ConcertDates: group.ConcertDates,
			Relation:     group.Relation,
		}
		data = append(data, artist)
	}
}
