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
	n := make(map[string]int)
	longestPaths := make(map[int]bool) 
	for i := 0; i < len(paths); i++ {
		if len(paths[i])==2{
			utils.Filter=append(utils.Filter, paths[i])

		}
		for j := 1; j < len(paths[i])-1; j++ {
			if k, exist := n[paths[i][j]]; exist {
				if len(paths[k]) < len(paths[i]) {
					longestPaths[k] = false 
					n[paths[i][j]] = i         
					longestPaths[i] = true 
				} else {
					if !longestPaths[k]{
						longestPaths[i] = true
						 n[paths[i][j]] = i  
					}else{
						longestPaths[i]=false
                         break
					}
				
					
				}
			} else {
				n[paths[i][j]] = i
				longestPaths[i] = true
			}
		}
	}

	for i, keep := range longestPaths {
		if keep {
			utils.Filter= append(utils.Filter, paths[i])
			
		}
	}
	

return utils.Filter
}
