package printage

import (
	"fmt"
	"strings"

	"lem-in/utils"
)

func Printage(colony *utils.AntFarm, paths [][]string) {
	ants := make([]utils.Ant, colony.NumAnts)
	antPositions := make(map[int]int)

	antIndex := 0
	for i := 0; i < colony.NumAnts; i++ {
		ants[i] = utils.Ant{ID: i + 1, Path: paths[antIndex]}
		antIndex++
		if antIndex >= len(paths) {
			antIndex = 0
		}
	}

	step := 0
	for {
		moved := false
		var moveStrs []string

		for i := 0; i < colony.NumAnts; i++ {
			if antPositions[i] < len(ants[i].Path)-1 {
				// Move to next position
				antPositions[i]++
				moveStrs = append(moveStrs, fmt.Sprintf("L%d-%s", ants[i].ID, ants[i].Path[antPositions[i]]))
				moved = true
			}
		}

		if !moved {
			break
		}

		fmt.Println(strings.Join(moveStrs, " "))
		step++
	}
}
