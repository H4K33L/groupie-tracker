package Groupie

import "strings"

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
