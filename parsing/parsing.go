package parsing

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem-in/utils"
)

func Parsing() *utils.Colony {
	if len(os.Args) != 2 {
		fmt.Println("Error: Please provide exactly one file name.")
		return nil
	}

	colony := &utils.Colony{
		Rooms: make(map[string]*utils.Room),
		Links: make(map[string][]string),
		Start: nil,
		End:   nil,
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error: Unable to read file:", err)
		return nil
	}

	lines := strings.Split(string(file), "\n")
	ant, err := strconv.Atoi(strings.TrimSpace(lines[0]))
	if err != nil || ant <= 0 {
		fmt.Println("Error: Invalid ant number")
		return nil
	}
	colony.NumAnts = ant

	for index := 1; index < len(lines); index++ {
		line := strings.TrimSpace(lines[index])

		if line == "" {
			continue
		}

		room := strings.Fields(line)
		if len(room) == 3 {
			if room[0][0] == 'L' {
				fmt.Println("Error: Room name cannot start with 'L'.")
				return nil
			}

			if _, exists := colony.Rooms[room[0]]; exists {
				fmt.Println("Error: Duplicate room name found:", room[0])
				return nil
			}

			colony.Rooms[room[0]] = &utils.Room{
				Name: room[0],
				X:    room[1],
				Y:    room[2],
			}
			continue
		}

		if line == "##start" {
			if index+1 >= len(lines) {
				fmt.Println("Error: Missing start room data.")
				return nil
			}
			startRoom := strings.Fields(strings.TrimSpace(lines[index+1]))
			if len(startRoom) != 3 {
				fmt.Println("Error: Invalid start room format.")
				return nil
			}
			colony.Start = &utils.Room{
				Name: startRoom[0],
				X:    startRoom[1],
				Y:    startRoom[2],
			}
			index++
			continue
		}

		if line == "##end" {
			if index+1 >= len(lines) {
				fmt.Println("Error: Missing end room data.")
				return nil
			}
			endRoom := strings.Fields(strings.TrimSpace(lines[index+1]))
			if len(endRoom) != 3 {
				fmt.Println("Error: Invalid end room format.")
				return nil
			}
			colony.End = &utils.Room{
				Name: endRoom[0],
				X:    endRoom[1],
				Y:    endRoom[2],
			}
			index++
			continue
		}

		link := strings.Split(line, "-")
		if len(link) == 2 {
			colony.Links[link[0]] = append(colony.Links[link[0]], link[1])
			colony.Links[link[1]] = append(colony.Links[link[1]], link[0])
		}
	}

	if colony.Start == nil {
		fmt.Println("Error: Start room is missing.")
		return nil
	}
	if colony.End == nil {
		fmt.Println("Error: End room is missing.")
		return nil
	}

	return colony
}
