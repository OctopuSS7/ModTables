package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	base, _ := strconv.Atoi(os.Args[2])
	operation := os.Args[1]
	if operation == "multiplicative" {
		exclude := []int{}
		generate_board(multiplicative, base, exclude)
	} else if operation == "additive" {
		exclude := []int{}
		generate_board(additive, base, exclude)
	} else if operation == "latin_square" {
		exclude := append(coprime(base), 0)
		generate_board(multiplicative, base, exclude)
	} else {
		fmt.Println("ERROR: Invalid Arguments")
	}
}
