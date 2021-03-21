package main

import (
	"fmt"
	"flag"
//	"strconv"
)

func parse_flags() (int, string) {
	var base int
	var mode string
	flag.IntVar(&base, "base", 10, "Which base should this table run in?")
	flag.StringVar(&mode, "mode", "i", "Multiplicative [m], Additive [a], or a Latin Square [l]?")
	flag.Parse()
	return base, mode
}

func main() {
	base, mode := parse_flags()
	if (mode == "m") {
		exclude := []int{}
		generate_board(multiplicative, base, exclude)
	} else if (mode == "a") {
		exclude := []int{}
		generate_board(additive, base, exclude)
	} else if (mode == "l") {
		exclude := append(coprime(base), 0)
		generate_board(multiplicative, base, exclude)
	} else if (mode == "i") {
		interactive()
	} else {
		fmt.Println("%d, %s", base, mode)
		fmt.Println("ERROR: Invalid Arguments")
	}
}
