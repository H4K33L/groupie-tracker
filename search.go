package Groupie

import (
	"fmt"
	"strings"
)

func GetGroupByID(groups []group, id int) (group, error) {
	for _, g := range groups {
		if int(g.ID) == id {
			return g, nil
		}
	}

	return group{}, fmt.Errorf("groupe avec l'ID %v non trouvé", id)
}

func GetGroupsByName(groups []group, partial string) []group {
	matchingGroups := []group{}
	for _, g := range groups {
		if strings.Contains(strings.ToLower(g.Name), strings.ToLower(partial)) {
			matchingGroups = append(matchingGroups, g)
		}
	}

	return matchingGroups
}

func GetGroupsByMember(groups []group, partialMemberName string) []group {
	matchingGroups := []group{}

	for _, g := range groups {
		for _, member := range g.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(partialMemberName)) {
				matchingGroups = append(matchingGroups, g)
				break
			}
		}
	}

	return matchingGroups
}
