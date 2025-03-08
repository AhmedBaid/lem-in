package utils

type Ant struct {
	ID   int
	Path []string
}

type AntFarm struct {
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
