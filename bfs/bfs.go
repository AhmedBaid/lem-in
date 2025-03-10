package graph

import (
	"lem-in/utils"
)

func FindPaths(colony *utils.AntFarm) [][]string {
	var paths [][]string
	queue := [][]string{{colony.Start.Name}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		lastRoom := path[len(path)-1]

		if lastRoom == colony.End.Name {
			paths = append(paths, path)
			continue
		}

		for _, neighbor := range colony.Links[lastRoom] {
			if !contains(path, neighbor) {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	return paths
}

// contains checks if a slice contains a specific item
func contains(slice []string, item string) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}
