package main

import (
	"fmt"
	"flag"
//	"strconv"
)

var base int
var mode string

func parseFlags() {
	flag.IntVar(&base, "base", 10, "Which base should this table run in?")
	flag.StringVar(&mode, "mode", "q", "Multiplicative, Additive, or a Latin Square?")
	flag.Parse()
}

func main() {
	parseFlags()
	if mode == "m" {
		exclude := []int{}
		generate_board(multiplicative, base, exclude)
	} else if mode == "a" {
		exclude := []int{}
		generate_board(additive, base, exclude)
	} else if mode == "l" {
		exclude := append(coprime(base), 0)
		generate_board(multiplicative, base, exclude)
	} else {
		fmt.Println("ERROR: Invalid Arguments")
	}
}
