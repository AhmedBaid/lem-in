package utils

type Room struct {
	Name string
	X,Y string
}
type AntFarm struct  {
	NumAnts int
	Start       *Room
	End         *Room
	Links       map[string][]string
	Rooms        map[string]*Room
}
