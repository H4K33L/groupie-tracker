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
    ID                  float64         `json:"id"`
    Img                 string          `json:"image"`
    Name                string          `json:"name"`
    Members             []string        `json:"members"`
    CreationDate        float64         `json:"creationDate"`
    FirstAlbum          string          `json:"firstAlbum"`
    Relation            string          `json:"relations"`
    Information         information
}
type information struct {
    Locations       []string            `json:"locations"`
    Dates           [][]string    `json:"dates"`
}
type temp struct {
    DatesLocation   map[string][]string  `json:"datesLocations"`
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

func GetLinkInfos(link string) (information, error) {
    matched, err := regexp.MatchString(`https:\/\/.*`, link)
    if err != nil {
        return information{}, err
    } else if matched {
        elements := information{}
        infos := temp{}

        response, err := http.Get(link)
        if err != nil {
            return information{}, err
        }

        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
            return information{}, err
        }

        err = json.Unmarshal(responseData, &infos)
        if err != nil {
            return information{}, err
        } 
        
        for clef := range infos.DatesLocation {
            elements.Locations = append(elements.Locations, clef)
            elements.Dates = append(elements.Dates, infos.DatesLocation[clef])
        }

        return elements, nil

    } else {
        return information{}, err
    }
}