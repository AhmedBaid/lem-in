package parsing

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem-in/utils"
)

func Parsing() *utils.AntFarm {
	colony := &utils.AntFarm{
		Start: &utils.Room{},
		End:   &utils.Room{},
		Rooms: make(map[string]*utils.Room),
		Links: make(map[string][]string),
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./lem-in [filename]")
		os.Exit(1)

	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return nil

	}

	line := strings.Split(string(file), "\n")
	nbrAnts, err := strconv.Atoi(line[0])
	if err != nil {
		fmt.Println("Error: ", err)
		return nil

	}

	if nbrAnts <= 0 {
		fmt.Println("Error: Number of ants must be greater than 0")
		return nil

	}

	colony.NumAnts = nbrAnts

	StartDup := false
	EndDup := false

	for i := 1; i < len(line); i++ {

		if line[i] == "" {
			continue
		}

		room := &utils.Room{
			Name: "",
			X:    "",
			Y:    "",
		}
		rooms := strings.Fields(line[i])

		if len(rooms) == 3 {
			if rooms[0][0] == 'L' || rooms[0][0] == '#' {
				fmt.Println("Error: Room name cannot start with 'L' or '#'")
				return nil
			}

			if _, exists := colony.Rooms[rooms[0]]; exists {

				fmt.Println("ERROR: invalid data format (duplicate room name)")
				return nil
			}

			for _, v := range colony.Rooms {
				if v.X == rooms[1] && v.Y == rooms[2] {
					fmt.Println("ERROR: invalid data format (duplicate coordinates)")
					return nil

				}
			}

			room.Name = rooms[0]
			room.X = rooms[1]
			room.Y = rooms[2]

			colony.Rooms[room.Name] = room

		}

		if strings.TrimSpace(line[i-1]) == "##start" {
			if StartDup {
				fmt.Println("ERROR: invalid data format (start or end  is depleted)")
				return nil
			}
			StartDup = true
			if len(rooms) == 3 {

				colony.Start.Name = rooms[0]
				colony.Start.X = rooms[1]
				colony.Start.Y = rooms[2]
			}

			continue
		}

		if strings.TrimSpace(line[i-1]) == "##end" {
			if EndDup {
				fmt.Println("ERROR: invalid data format (start or end  is depleted)")
				return nil
			}
			EndDup = true

			if len(rooms) == 3 {

				colony.End.Name = rooms[0]
				colony.End.X = rooms[1]
				colony.End.Y = rooms[2]
			}
			continue
		}

		link := strings.Split(line[i], "-")

		if len(link) == 2 {

			colony.Links[link[0]] = append(colony.Links[link[0]], link[1])
			colony.Links[link[1]] = append(colony.Links[link[1]], link[0])

			continue

		}

	}

	if colony.Start.Name == "" || colony.End.Name == "" || len(colony.Rooms) == 0 {
		fmt.Println("ERROR: invalid data format (missing start, end or rooms)")
		return nil
	}

	for  v := range colony.Links {
		if _, exists := colony.Rooms[v]; !exists {
			fmt.Println("ERROR: invalid data format (missing links)")
			return nil

		}
	}



	return colony
}
