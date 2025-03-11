package main

import (
	"fmt"

	graph "lem-in/bfs"
	"lem-in/parsing"
)

func main() {
	colony := parsing.Parsing()
resu := graph.Bfs(colony)	

/*
		 	if colony == nil {
				os.Exit(1)
			}

			paths := graph.Bfs(colony)
			if len(paths) == 0 {
				utils.PrintError(fmt.Errorf("NO VALID PATHS FOUND"))
				return
			}

			printage.Printage(colony, paths)
	*/
}
