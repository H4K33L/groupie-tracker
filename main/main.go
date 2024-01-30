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
	fmt.Println(information)
	groups, err := Groupie.GetArtist(information.Artists)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(groups[0])
}