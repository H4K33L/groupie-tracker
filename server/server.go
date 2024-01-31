package main

import (
	"Groupie"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type api struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}
type group struct {
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

var data group

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("../template/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

//var tmpl = template.Must(template.ParseFiles())

func main() {

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	start()
	http.HandleFunc("/index", index)

	fmt.Printf("Starting server to test (listen 8080)\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	RenderTemplates(w, "index")
	t, err := template.ParseFiles("../template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

func GetApi(link string) (api, error) {
	apiLink := api{"https://groupietrackers.herokuapp.com/api/artists", "https://groupietrackers.herokuapp.com/api/locations", "https://groupietrackers.herokuapp.com/api/dates", "https://groupietrackers.herokuapp.com/api/relation"}
	response, err := http.Get(link)
	if err != nil {
		return apiLink, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return apiLink, err
	}

	err = json.Unmarshal(responseData, &apiLink)
	if err != nil {
		return apiLink, err
	}

	return apiLink, nil
}

func GetArtist(link string) ([]group, error) {
	groups := []group{}
	response, err := http.Get(link)
	if err != nil {
		return groups, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return groups, err
	}

	err = json.Unmarshal(responseData, &groups)
	if err != nil {
		return groups, err
	}

	return groups, nil
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

	data.ID = (groups[0].ID)
	data.Img = (groups[0].Img)
	data.Name = (groups[0].Name)
	data.Members = (groups[0].Members)
	data.CreationDate = (groups[0].CreationDate)
	data.FirstAlbum = (groups[0].FirstAlbum)
	data.Locations = (groups[0].Locations)
	data.ConcertDates = (groups[0].ConcertDates)
	data.Relation = (groups[0].Relation)

}
