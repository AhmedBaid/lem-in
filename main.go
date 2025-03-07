package main

import (
	"fmt"
	"lem-in/bfs"
	"lem-in/parsing"
	"lem-in/printage"
	"lem-in/utils"
	"os"
)

func main() {
	colony := parsing.Parsing()
	if colony == nil {
		os.Exit(1)
	}


	paths := graph.FindPaths(colony)
	if len(paths) == 0 {
		utils.PrintError(fmt.Errorf("No valid paths found"))
		return
	}

	printage.MoveAnts(colony, paths)
}
