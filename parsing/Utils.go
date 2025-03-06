package parsing

type Room struct {
	name string
	x,y string
}
type Colony struct {
	NumAnts int
	start       *Room
	end         *Room
	links       map[string][]string
	rooms        map[string]*Room
}
