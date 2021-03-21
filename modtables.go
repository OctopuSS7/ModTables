package main

import "fmt"

func generate_board(op operator, base int, exclude []int) {
	board := [][]int{}
	for i := 0; i < base; i++ {
		row := []int{}
		for j := 0; j < base; j++ {
			if !(contains(exclude, i) || contains(exclude, j)) {
				row = append(row, op(i, j, base))
			}
		}
		if len(row) > 0 {
			board = append(board, row)
		}
	}
	render_board(board)
}

func render_board(board [][]int) {
	for _, x := range board {
		fmt.Print("| ")
		for _, y := range x {
			fmt.Print(y)
			fmt.Print(" | ")
		}
		fmt.Println("")
	}
}
