package Groupie

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
	"regexp"
)

// all the structures used to represent api information
type api struct {
	Artists 	string `json:"artists"`
	Locations 	string `json:"locations"`
	Dates 		string `json:"dates"`
	Relation 	string `json:"relation"`
}
type group struct {
    ID                  float64 `json:"id"`
    Img                 string  `json:"image"`
    Name                string  `json:"name"`
    Members             []string  `json:"members"`
    CreationDate        float64 `json:"creationDate"`
    FirstAlbum          string  `json:"firstAlbum"`
    Locations           string  `json:"locations"`
    ConcertDates        string  `json:"concertDates"`
    Relation            string  `json:"relations"`
    ListeLocations      []string
    ListConcertDates    []string
    ListRelation        map[string]interface{}
}
type information struct {
    ID              float64  `json:"id"`
    Locations       []string `json:"locations"`
    Dates           string `json:"dates"`
    DatesLocations  map[string]interface{} `json:"datesLocations"`
}
type infoSup struct {
    Dates           []string `json:"dates"`
}

func GetApi(link string) (api,error) {
    apiLink := api{}
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

func GetLinkElem(link string) ([]string, map[string]interface{}, error) {
    matched, err := regexp.MatchString(`https:\/\/.*`, link)
    if err != nil {
        return []string{}, map[string]interface{}{}, err
    } else if matched {
        elements := information{}
        response, err := http.Get(link)
        if err != nil {
            return []string{}, map[string]interface{}{}, err
        }

        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
            return []string{}, map[string]interface{}{}, err
        }
        err = json.Unmarshal(responseData, &elements)
        if err != nil {
            element := infoSup{}
            err = json.Unmarshal(responseData, &element)
            if err != nil {
                return []string{}, map[string]interface{}{}, err
            }
            return element.Dates, map[string]interface{}{}, err
        }
        if len(elements.Locations) != 0 {
            return elements.Locations, map[string]interface{}{}, nil
        } else {
            return []string{}, elements.DatesLocations, nil
        }
    } else {
        return []string{}, map[string]interface{}{}, nil
    }
}