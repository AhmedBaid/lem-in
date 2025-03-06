package utils

type Rooms struct {
	Name string
	X, Y string
}
type IntFarm struct {
	Number_ants int
	Start       *Rooms
	End         *Rooms
	Link        map[string][]string
	Room        map[string]*Rooms
}
