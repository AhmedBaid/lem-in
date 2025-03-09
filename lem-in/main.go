package main

import (
	"fmt"
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
		fmt.Println("No valid paths found.")
		os.Exit(1)
	}

	printage.Printage(colony, paths)
}
