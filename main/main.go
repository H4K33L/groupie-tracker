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
	groups[1].Information, err = Groupie.GetLinkInfos(groups[1].Relation)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(groups[1])
	fmt.Println(groups[1].Information.Locations[2], groups[1].Information.Dates[2][0])
}