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

	var level int
	var name string
	var member string

	fmt.Printf("entrez l'ID : ")

	_, err = fmt.Scan(&level)
	if err != nil {
		fmt.Println(err)
		return
	}

	group, err := Groupie.GetGroupByID(groups, level)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(group)

	fmt.Print("Entrez une partie ou le nom complet du groupe: ")
	_, err = fmt.Scan(&name)

	if name == "\n" {
		fmt.Println(group)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	groups2 := Groupie.GetGroupsByName(groups, name)

	if len(groups2) == 0 {
		fmt.Printf("Aucun groupe correspondant au nom '%s' n'a été trouvé.\n", name)
	} else {
		fmt.Printf("Groupes correspondants:\n")
		for _, g := range groups2 {
			fmt.Printf("%+v\n", g)
		}
	}

	fmt.Print("Entrez une partie ou le nom complet d'un membre du groupe :")
	_, err = fmt.Scan(&member)
	if err != nil {
		fmt.Println(err)
		return
	}

	groups3 := Groupie.GetGroupsByMember(groups, member)

	if len(groups3) == 0 {
		fmt.Printf("Aucun groupe correspondant au nom '%s' n'a été trouvé.\n", name)
	} else {
		fmt.Printf("Groupes correspondants:\n")
		for _, g := range groups3 {
			fmt.Printf("%+v\n", g)
		}
	}

}
