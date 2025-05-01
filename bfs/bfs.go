package graph

import (
	"fmt"
	"runtime"

	"lem-in/utils"
)

/*
	 func BFS(graph map[string][]string, start, end string) []string {
		queue := [][]string{{start}}

		for len(queue) > 0 {
			path := queue[0]
			queue = queue[1:]
			last := path[len(path)-1]

			if last == end {
				return path
			}

			for _, neighbor := range graph[last] {
				if !contains(path, neighbor){

					newPath := append([]string{}, path...)
					newPath = append(newPath, neighbor)
					queue = append(queue, newPath)
				}
			}
		}

		return nil
	}

	func FindPaths(colony *utils.AntFarm) [][]string {
		var allPaths [][]string

		for _, startNeighbor := range colony.Links[colony.Start.Name] {
			// Loop on neighbors of End
			for _, endNeighbor := range colony.Links[colony.End.Name] {
				fmt.Println("start", startNeighbor)
				fmt.Println("end", endNeighbor)

				// Get path from startNeighbor to endNeighbor
				pathStartToEnd := BFS(colony.Links, startNeighbor, endNeighbor)

fmt.Println("paaa", pathStartToEnd)

				// If a path was found, complete it by adding Start and End
				if pathStartToEnd != nil {
					// Add Start at the beginning and End at the end
					fullPath := append([]string{colony.Start.Name}, pathStartToEnd...)
					fullPath = append(fullPath, colony.End.Name)
					allPaths = append(allPaths, fullPath)
				}
			}
		}

		return allPaths
	}
*/
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

				// نحدّدو حد أقصى مثلاً 100MB
				highMemory := memStats.Alloc > 100*1024*1024

				if highMemory {
					// فاش تكون الذاكرة مستهلكة بزاف
					if !contains(path, next) && !visited[next] {
						newPath := append([]string{}, path...)
						newPath = append(newPath, next)
						queue = append(queue, newPath)
						visited[next] = true
					}
				} else {
					// فاش تكون الذاكرة ماشي مشكل
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
	fmt.Println(utils.Filter)
	return utils.Filter
}
