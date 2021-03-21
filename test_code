func render_board(op operator, base int, exclude []int) {
	for i := 0; i < base; i++ {
		if !(contains(exclude, i)) {
			fmt.Print("| ")
		}
		for j := 0; j < base; j++ {
			if !(contains(exclude, i) || contains(exclude, j)) {
				fmt.Print(op(i, j, base))
				fmt.Print(" | ")
			} //else {fmt.Print("  ")}
		}
		if !(contains(exclude, i)) {
			fmt.Println("")
			fmt.Println("")
		}
	}
}

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
