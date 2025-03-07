package graph

import "lem-in/utils"

func FindPaths(colony *utils.Colony) [][]string {
	queue := [][]string{{colony.Start.Name}}
	var paths [][]string
	visited := map[string]bool{colony.Start.Name: true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastRoom := path[len(path)-1]

		if lastRoom == colony.End.Name {
			paths = append(paths, path)
			continue
		}

		for _, neighbor := range colony.Links[lastRoom] {
			if !visited[neighbor] {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
				visited[neighbor] = true
			}
		}
	}

	return paths
}
