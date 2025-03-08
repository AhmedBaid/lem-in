package main

import (
	"os"

	graph "lem-in/bfs"
	"lem-in/parsing"
	"lem-in/printage"
)

func main() {
	colony := parsing.Parsing()
	if colony == nil {
		os.Exit(1)
	}

	paths := graph.FindPaths(colony)
	if len(paths) == 0 {
		// utils.PrintError(fmt.Errorf("NO VALID PATHS FOUND"))
		return
	}

	printage.Printage(colony, paths)
}
