package parsing_test

import (
	"fmt"
	"os"
	"strings"

	"lem-in/utils"
)

func Parse() {
	x, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error", err)
	}
	lin := strings.Split(string(x), "\n")
	if err != nil {
		fmt.Println("Error", err)
	}

	// Number_ants, err := strconv.Atoi(lin[0])
	var Rooms utils.Rooms
	for i := 0; i < len(lin); i++ {
		if lin[i] == "##start" && lin[i+1] != "##end" {
			cord := strings.Fields(lin[i+1])

			Rooms = utils.Rooms{
				Name: string(cord[0]),
				X:    string(cord[1]),
				Y:    string(cord[2]),
			}


		}else if lin[i] == "##end"  {
	}
	fmt.Println(Rooms)
	fmt.Println(lin[1])
}
