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
	var currentPaths [][]string        
	usedNodes := make(map[string]bool) 
	var backtrack func(int)
	backtrack = func(start int) {
		if len(currentPaths) > len(utils.Filter) {
			utils.Filter = make([][]string, len(currentPaths))
			copy(utils.Filter, currentPaths)
		}

		for i := start; i < len(paths); i++ {
			path := paths[i]
			keepPath := true 

			for _, node := range path[1 : len(path)-1] {
				if usedNodes[node] {
					keepPath = false 
					break
				}
			}

			if keepPath {
				currentPaths = append(currentPaths, path)
				for _, node := range path[1 : len(path)-1] {
					usedNodes[node] = true
				}

				backtrack(i + 1)

				currentPaths = currentPaths[:len(currentPaths)-1]
				for _, node := range path[1 : len(path)-1] {
					delete(usedNodes, node)
				}
			}
		}
	}

	backtrack(0)
	return utils.Filter
}
