package printage

import (
	"fmt"
	"strings"
	"lem-in/utils"
)

func Printage(colony *utils.AntFarm, paths [][]string) {
	// Create a slice of ants and assign each ant to a path.
	ants := make([]utils.Ant, colony.NumAnts)
	antPositions := make(map[int]int)

	// Distribute ants across paths, prioritizing shorter paths
	antIndex := 0
	for i := 0; i < colony.NumAnts; i++ {
		ants[i] = utils.Ant{ID: i + 1, Path: paths[antIndex]}
		antIndex++
		if antIndex >= len(paths) {
			antIndex = 0
		}
	}

	step := 0
	// Start the simulation of ant movement
	for {
		moved := false
		var moveStrs []string

		// Go through each ant and move them to the next room if possible
		for i := 0; i < colony.NumAnts; i++ {
			if antPositions[i] < len(ants[i].Path)-1 {
				// Move the ant to the next room
				antPositions[i]++
				moveStrs = append(moveStrs, fmt.Sprintf("L%d-%s", ants[i].ID, ants[i].Path[antPositions[i]]))
				moved = true
			}
		}

		// If no ants moved, end the simulation
		if !moved {
			break
		}

		// Print the current step's movements of all ants
		fmt.Println(strings.Join(moveStrs, " "))
		step++
	}
}
