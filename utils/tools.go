package utils

type AntFarm struct {
	Start *Room
	End   *Room
	Rooms map[string]*Room
	Links map[string][]string
}

type Room struct {
	Name string
	X    string
	Y    string
}

type Node struct {
	Value     string
	Neighbors []string
}

var Paths [][]string

var Filter [][]string

var Ants int
