package graph

import (
	"runtime"

	"lem-in/utils"
)

func FindPaths(colony *utils.AntFarm) [][]string {
	paths := [][]string{}

	for _, neighbor := range colony.Links[colony.Start.Name] {
		visited := map[string]bool{colony.Start.Name: true, neighbor: true}
		queue := [][]string{{colony.Start.Name, neighbor}}

		for len(queue) > 0 {
			path := queue[0]
			queue = queue[1:]

			lastRoom := path[len(path)-1]
			if lastRoom == colony.End.Name {
				paths = append(paths, path)
				continue
			}

			for _, next := range colony.Links[lastRoom] {
				var memStats runtime.MemStats
				runtime.ReadMemStats(&memStats)

				highMemory := memStats.Alloc > 100*1024*1024

				if highMemory {
					if !contains(path, next) && !visited[next] {
						newPath := append([]string{}, path...)
						newPath = append(newPath, next)
						queue = append(queue, newPath)
						visited[next] = true
					}
				} else {
					if !contains(path, next) {
						newPath := append([]string{}, path...)
						newPath = append(newPath, next)
						queue = append(queue, newPath)
						visited[next] = true
					}
				}
			}
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
