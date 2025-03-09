package utils

type AntFarm struct {
	Start  *Room
	End    *Room
	Rooms  map[string]*Room
	Links  map[string][]string
	NumAnts int
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

type Ant struct {
	ID   int
	Path []string
}
