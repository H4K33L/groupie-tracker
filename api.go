package Groupie

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
	"regexp"
	"errors"
)

// all the structures used to represent api information
type api struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
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
    DatesLocation   map[string][]string  `json:"datesLocations"`
}

var (
    GetApiError = errors.New("There is an error during the API calling, OH SNAP !")
    ConvertToStringError = errors.New("There is an error during the convertion to string, SOOOOOOOOOOOOOOOOOO BAD !")
	UnmarshalError = errors.New("The json probably doesnt match with the struct.....TRY AGAIN !")
    RegexpError = errors.New("An error append in the Regexp rule, IT WORK ON MY PC !")
    RegexpNoMatch = errors.New("Sorry but this isnt an URL, TOU CAN DO IT BETTER !")
)

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
    matched, err := regexp.MatchString(`https:\/\/.*`, link)
    if err != nil {
        return api{}, RegexpError
    } else if matched {
        apiLink := api{}
        response, err := http.Get(link)
        if err != nil {
            return apiLink, GetApiError
        }

        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
            return apiLink, ConvertToStringError
        }

	    err = json.Unmarshal(responseData, &apiLink)
	    if err != nil {
            return apiLink, UnmarshalError
        }
	
        return apiLink,nil
    } else {
        return api{}, RegexpNoMatch
    }
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
    matched, err := regexp.MatchString(`https:\/\/.*`, link)
    if err != nil {
        return []group{}, RegexpError
    } else if matched {
        groups := []group{}
        response, err := http.Get(link)
        if err != nil {
            return groups, GetApiError
        }

        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
            return groups, ConvertToStringError
        }

	    err = json.Unmarshal(responseData, &groups)
	    if err != nil {
            return groups, UnmarshalError
        }
	
        return groups,nil
    } else {
        return []group{}, RegexpNoMatch
    }
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
        return information{}, RegexpError
    } else if matched {
        elements := information{}

        response, err := http.Get(link)
        if err != nil {
            return information{}, GetApiError
        }

        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
            return information{}, ConvertToStringError
        }

        err = json.Unmarshal(responseData, &elements)
        if err != nil {
            return information{}, UnmarshalError
        } 

        return elements, nil

    } else {
        return information{}, RegexpNoMatch
    }
}


func Begin(link string) ([]group,error) {
    /*
    The Begin function initillize the group list
    ---------------------------------------------
    input : the link to the first api.
    output : the list of group
    ---------------------------------------------
    The function also return the posible error
    produce by the program.
    */
	information, err := GetApi(link)
	if err != nil {
		return []group{}, err
	}
	groups, err := GetArtist(information.Artists)
	if err != nil {
		return []group{}, err
	}
	return groups, nil
}

func LoadArtist(link, letter string) (group, error) {
    /*
    The Load Artist func is used to get information about one group whith is ID.
    -----------------------------------------------------------------------------
    input : the api link and the id of the group
    output : the group struct who contain the group information
    -----------------------------------------------------------------------------
    Also the function return the posible error apening during the execution.
    */
    information := api{}
    artist := group{}

    information, err := GetApi(link)
	if err != nil {
		return group{}, err
	}

    response, err := http.Get(information.Artists+"/"+letter)
    if err != nil {
        return group{}, GetApiError
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return group{}, ConvertToStringError
    }

    err = json.Unmarshal(responseData, &artist)
    if err != nil {
        return group{}, UnmarshalError
    } 

    return artist, nil
}