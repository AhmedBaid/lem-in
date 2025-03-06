package utils

type Rooms struct {
	Name string
	X, Y string
}
type Khaliya struct {
	Number_ants int
	Start       *Rooms
	End         *Rooms
	Link        map[string][]string
}
