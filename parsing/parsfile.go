package parsing

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem-in/utils"
)

func Parsing() *utils.AntFarm {
	// Create a new AntFarm struct
	colony := &utils.AntFarm{
		Start: &utils.Room{},
		End:   &utils.Room{},
		Rooms: make(map[string]*utils.Room),
		Links: make(map[string][]string),
	}
	// check if the number of arguments is equal to 2
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./lem-in [filename]")
		return nil

	}
	// read the file
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return nil

	}
	// split the file into lines
	line := strings.Split(string(file), "\n")
	nbrAnts, err := strconv.Atoi(line[0])
	if err != nil {
		fmt.Println("Error: ", err)
		return nil

	}
	// check if the number of ants is greater than 0
	if nbrAnts <= 0 {
		fmt.Println("Error: Number of ants must be greater than 0")
		return nil

	}
	// assign the number of ants to the colony
	colony.NumAnts = nbrAnts

	StartDup := false
	EndDup := false

	for i := 1; i < len(line); i++ {

		if line[i] == "" {
			continue
		}
		// create a new room
		room := &utils.Room{
			Name: "",
			X:    "",
			Y:    "",
		}
		rooms := strings.Fields(line[i])
		// check if the number of rooms is equal to 3
		if len(rooms) == 3 {
			//  check if starts with 'L' or '#'
			if rooms[0][0] == 'L' || rooms[0][0] == '#' {
				fmt.Println("Error: Room name cannot start with 'L' or '#'")
				return nil
			}

			// check if the room name is already in the map
			if _, exists := colony.Rooms[rooms[0]]; exists {

				fmt.Println("ERROR: invalid data format (duplicate room name)")
				return nil
			}
			// check if the coordinates are duplicates
			for _, v := range colony.Rooms {
				if v.X == rooms[1] && v.Y == rooms[2] {
					fmt.Println("ERROR: invalid data format (duplicate coordinates)")
					return nil

				}
			}
			// assign the room name, x and y coordinates
			room.Name = rooms[0]
			room.X = rooms[1]
			room.Y = rooms[2]

			colony.Rooms[room.Name] = room

		}

		// check if the line starts with '##start'
		if strings.TrimSpace(line[i-1]) == "##start" {
			// check if the start is duplicated
			if StartDup {
				fmt.Println("ERROR: invalid data format (start or end  is depleted)")
				return nil
			}
			StartDup = true
			// add the start room to the colony
			if len(rooms) == 3 {

				colony.Start.Name = rooms[0]
				colony.Start.X = rooms[1]
				colony.Start.Y = rooms[2]
			}

			continue
		}
		// check if the line starts with '##end'
		if strings.TrimSpace(line[i-1]) == "##end" {
			// check if the end is duplicated
			if EndDup {
				fmt.Println("ERROR: invalid data format (start or end  is depleted)")
				return nil
			}
			EndDup = true
			// add the end room to the colony
			if len(rooms) == 3 {

				colony.End.Name = rooms[0]
				colony.End.X = rooms[1]
				colony.End.Y = rooms[2]
			}
			continue
		}

		link := strings.Split(line[i], "-")
		// add the links to the colony
		if len(link) == 2 {

			colony.Links[link[0]] = append(colony.Links[link[0]], link[1])
			colony.Links[link[1]] = append(colony.Links[link[1]], link[0])

			continue

		}

	}
	// check if the start, end or rooms are missing
	if colony.Start.Name == "" || colony.End.Name == "" || len(colony.Rooms) == 0 {
		fmt.Println("ERROR: invalid data format (missing start, end or rooms)")
		return nil
	}
	// check if the links are missing (if have a link of unkonwn room)
	for v := range colony.Links {
		if _, exists := colony.Rooms[v]; !exists {
			fmt.Println("ERROR: invalid data format (missing links)")
			return nil

		}
	}

	return colony
}
