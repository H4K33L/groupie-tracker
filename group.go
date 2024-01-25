package Groupie

import "fmt"

func GetGroupByID(groups []group, id int) (group, error) {
	for _, g := range groups {
		if int(g.ID) == id {
			return g, nil
		}
	}

	return group{}, fmt.Errorf("groupe avec l'ID %v non trouvé", id)
}
