package Groupie

import "strings"

func GetGroupsByName(groups []group, partial string) []group {
	matchingGroups := []group{}
	for _, g := range groups {
		if strings.Contains(strings.ToLower(g.Name), strings.ToLower(partial)) {
			matchingGroups = append(matchingGroups, g)
		}
	}

	return matchingGroups
}
