package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "encoding/json"
)

type api struct {
	artists string
	locations string
	dates string
	relation string
}

var elemn map[string]interface{}
var link api

func main() {
    response, err := http.Get("https://groupietrackers.herokuapp.com/api")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

	err = json.Unmarshal(responseData, &elemn)
	if err != nil {
        log.Fatal(err)
    }
	
    link.artists = elemn["artists"].(string)
	link.locations = elemn["locations"].(string)
	link.dates = elemn["dates"].(string)
	link.relation = elemn["relation"].(string)
    fmt.Println(link)
}