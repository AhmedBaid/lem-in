package printage

import (
	"fmt"
	"lem-in/utils"
	"strings"
)

func MoveAnts(colony *utils.Colony, paths [][]string) {
	ants := make([]utils.Ant, colony.NumAnts)
	for i := 0; i < colony.NumAnts; i++ {
		ants[i] = utils.Ant{ID: i + 1, Path: paths[i%len(paths)]}
	}

	step := 0
	for {
		moved := false
		var moveStrs []string

		for i := range ants {
			if step < len(ants[i].Path)-1 {
				moveStrs = append(moveStrs, fmt.Sprintf("L%d-%s", ants[i].ID, ants[i].Path[step+1]))
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
