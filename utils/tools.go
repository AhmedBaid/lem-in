package utils

import "fmt"

func PrintError(err error) {
	fmt.Println("ERROR:", err)
}

type Ant struct {
	ID   int
	Path []string
}

type Colony struct {
	NumAnts int
	Rooms   map[string]*Room
	Links   map[string][]string
	Start   *Room
	End     *Room
}

type Room struct {
	Name string
	X, Y string
}
