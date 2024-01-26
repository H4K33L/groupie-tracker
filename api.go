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
    Dates           [][]string          `json:"dates"`
}
type temp struct {
    DatesLocation   map[string][]string  `json:"datesLocations"`
}


func GetApi(link string) (api,error) {
    /*
    The GetApi func is used to get the first api information,
    the others api link, the function get the api response and 
    return all the link send by the api.
    -----------------------------------------------------------
    input : the URL to the main API
    output : all the information send by the API
    -----------------------------------------------------------
    The function return also all possible error case posibly
    generated during the execution.
    */
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
    /*
    The GetArtist function take the link to the API who 
    return all information about groups and return it 
    as a list.
    ----------------------------------------------------
    input : the URL to the API
    output : the list who contain all groups information
    ----------------------------------------------------
    The function also return error case posibly generated
    during the proces.
    */
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
    /*
    The GetLinkInfos function is used to get all time and localisation about concerts
    of one group. The information is two list, one of string and another whith list 
    of string.
    ----------------------------------------------------------------------------------
    input : an URL to an API
    output : all the information the api send
    ----------------------------------------------------------------------------------
    The function also return error case who probably can append during the proces.
    */
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