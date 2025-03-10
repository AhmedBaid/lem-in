package graph

import (
	"lem-in/utils"
)

func FindPaths(colony *utils.AntFarm) [][]string {
	queue := [][]string{{colony.Start.Name}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		lastRoom := path[len(path)-1]

		if lastRoom == colony.End.Name {
			utils.Paths = append(utils.Paths, path)
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

	return utils.Paths
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

func FindDisjointPaths(paths [][]string, colony *utils.AntFarm) [][]string {
	usedRooms := make(map[string]bool)

	for _, path := range paths {
		conflict := false

		for _, room := range path {
			if usedRooms[room] && room != colony.End.Name {
				conflict = true
				break
			}
		}

		if !conflict {
			utils.Filter = append(utils.Filter, path)
			for _, room := range path {
				usedRooms[room] = true
			}
		}
	}

	return utils.Filter
}
