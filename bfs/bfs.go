package graph

import (
	"lem-in/utils"
)
func FindShortestPath(start, end string, links map[string][]string, forbidden map[string]bool) []string {
	type Node struct {
		Name string
		Path []string
	}

	visited := make(map[string]bool)
	queue := []Node{{Name: start, Path: []string{start}}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Name == end {
			return current.Path
		}

		if visited[current.Name] {
			continue
		}
		visited[current.Name] = true

		for _, neighbor := range links[current.Name] {
			if !visited[neighbor] && !contains(current.Path, neighbor) && !forbidden[neighbor] {
				newPath := append([]string{}, current.Path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, Node{Name: neighbor, Path: newPath})
			}
		}
	}

	return nil
}

func FindPaths(colony *utils.AntFarm) [][]string {
	var paths [][]string

	for _, startNeighbor := range colony.Links[colony.Start.Name] {
		forbidden := map[string]bool{
			colony.Start.Name: true,
		}

		path := FindShortestPath(startNeighbor, colony.End.Name, colony.Links, forbidden)
		if path != nil {
			fullPath := append([]string{colony.Start.Name}, path...)
			paths = append(paths, fullPath)
		}
	}

	return paths
}


func contains(path []string, room string) bool {
	for _, r := range path {
		if r == room {
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
