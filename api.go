package Groupie

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
)

// all the structures used to represent api information
type api struct {
	Artists 	string `json:"artists"`
	Locations 	string `json:"locations"`
	Dates 		string `json:"dates"`
	Relation 	string `json:"relation"`
}
type group struct {
    ID              float64 `json:"id"`
    Img             string  `json:"image"`
    Name            string  `json:"name"`
    Members         []string  `json:"members"`
    CreationDate    float64 `json:"creationDate"`
    FirstAlbum      string  `json:"firstAlbum"`
    Locations       string  `json:"locations"`
    ConcertDates    string  `json:"concertDates"`
    Relation        string  `json:"relation"`
}

func GetApi(link string) (api,error) {
    apiLink := api{"https://groupietrackers.herokuapp.com/api"}
    response, err := http.Get(link)
    if err != nil {
        return apiLink,err
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return apiLink,err
    }

	err = json.Unmarshal(responseData, &apiLink)
	if err != nil {
        return apiLink,err
    }
	
    return apiLink,nil
}

func GetArtist(link string) ([]group,error) {
    groups := []group{}
    response, err := http.Get(link)
    if err != nil {
        return groups,err
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return groups,err
    }

	err = json.Unmarshal(responseData, &groups)
	if err != nil {
        return groups,err
    }
	
    return groups,nil
}