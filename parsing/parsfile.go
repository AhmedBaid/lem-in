package parsing

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parsing() (*Colony, string) {
	if len(os.Args) != 2 {
		fmt.Println("Error: Please provide exactly one file name.")
		return nil, ""
	}

	colony := &Colony{
		rooms: make(map[string]*Room),
		links: make(map[string][]string),
		start: nil,
		end:   nil,
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error: Unable to read file:", err)
		return nil, ""
	}

	lines := strings.Split(string(file), "\n")

	ant, err := strconv.Atoi(strings.TrimSpace(lines[0]))
	if err != nil || ant <= 0 {
		fmt.Println("Error: Invalid ant number")
		return nil, ""
	}
	colony.NumAnts = ant

	for r := 1; r < len(lines); r++ {
		line := strings.TrimSpace(lines[r])

		if line == "" {
			continue
		}

		room := strings.Fields(line)
		if len(room) == 3 {
			if room[0][0] == 'L' {
				fmt.Println("Error: Room name cannot start with 'L'.")
				return nil, ""
			}

			if _, exists := colony.rooms[room[0]]; exists {
				fmt.Println("Error: Duplicate room name found:", room[0])
				return nil, ""
			}

			for _, existingRoom := range colony.rooms {
				if room[1] == existingRoom.x && room[2] == existingRoom.y {
					fmt.Println("Error: Duplicate room coordinates found:", room[1], room[2])
					return nil, ""
				}
			}

			colony.rooms[room[0]] = &Room{
				name: room[0],
				x:    room[1],
				y:    room[2],
			}
			continue
		}

		if line == "##start" {
			if r+1 >= len(lines) {
				fmt.Println("Error: Missing start room data.")
				return nil, ""
			}
			startRoom := strings.Fields(strings.TrimSpace(lines[r+1]))
			if len(startRoom) != 3 {
				fmt.Println("Error: Invalid start room format.")
				return nil, ""
			}
			colony.start = &Room{
				name: startRoom[0],
				x:    startRoom[1],
				y:    startRoom[2],
			}
			r++
			continue
		}

		if line == "##end" {
			if r+1 >= len(lines) {
				fmt.Println("Error: Missing end room data.")
				return nil, ""
			}
			endRoom := strings.Fields(strings.TrimSpace(lines[r+1]))
			if len(endRoom) != 3 {
				fmt.Println("Error: Invalid end room format.")
				return nil, ""
			}
			colony.end = &Room{
				name: endRoom[0],
				x:    endRoom[1],
				y:    endRoom[2],
			}
			r++
			continue
		}

		link := strings.Split(line, "-")
		if len(link) == 2 {
			colony.links[link[0]] = append(colony.links[link[0]], link[1])
			colony.links[link[1]] = append(colony.links[link[1]], link[0])
		}
	}

	if colony.start == nil {
		fmt.Println("Error: Start room is missing.")
		return nil, ""
	}
	if colony.end == nil {
		fmt.Println("Error: End room is missing.")
		return nil, ""
	}

	return colony, string(file)
}