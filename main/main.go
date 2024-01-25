package main

import (
	"fmt"

	"Groupie"
)

func main() {
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
	groups[1].ListConcertDates,_, err = Groupie.GetLinkElem(groups[1].ConcertDates)
	if err != nil {
		fmt.Println(err)
		return
	}
	_,groups[1].ListRelation, err = Groupie.GetLinkElem(groups[1].Relation)
	if err != nil {
		fmt.Println(err)
		return
	}
	groups[1].ListeLocations,_, err = Groupie.GetLinkElem(groups[1].Locations)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(groups[1])
}