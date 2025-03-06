package parsing_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse() {
	x, err := os.ReadFile(os.Args[1])
	if err != nil {
	}
	lin := strings.Split(string(x), "\n")
	Number_ants, err := strconv.Atoi(lin[0])
	if err != nil {
		fmt.Println("efef")
	}
	fmt.Println(Number_ants)
}
