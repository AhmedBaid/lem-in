package main

import (
	"fmt"

	parsing_test "lem-in/parsing"
)

func main() {
	pars, file := parsing_test.Parsing()
	fmt.Println(file)
	fmt.Println(pars)
}
