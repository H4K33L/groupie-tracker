package Groupie

import (
	"fmt"
	"strings"
)

// function that allow us to search one artist by ID
func GetGroupByID(groups []group, id int) ([]group, error) {
	for _, g := range groups {
		if int(g.ID) == id {
			return []group{g}, nil
		}
	}

	return []group{}, fmt.Errorf("groupe avec l'ID %v non trouvé", id)
}

// function that allow us to search one artist by his partial or full name
func GetGroupsByName(groups []group, partial string) []group {
	matchingGroups := []group{}
	for _, g := range groups {
		if strings.Contains(strings.ToLower(g.Name), strings.ToLower(partial)) { //all strings to lowercase
			matchingGroups = append(matchingGroups, g)
		}
	}

	return matchingGroups
}

// function that allow us to search one artist by a group member
func GetGroupsByMember(groups []group, partialMemberName string) []group {
	matchingGroups := []group{}

	for _, g := range groups {
		for _, member := range g.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(partialMemberName)) { //all strings to lowercase
				matchingGroups = append(matchingGroups, g)
				break
			}
		}
	}

	return matchingGroups
}
