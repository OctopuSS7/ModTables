package main

import "fmt"

func interactive() {
  var mode string
  var base int
  for {
    fmt.Printf("(Base) [-->]: ")
    fmt.Scanf("%d\n", &base)
    fmt.Printf("(Mode) [-->]: ")
    fmt.Scanf("%s\n", &mode)
    if (mode == "m") {
      exclude := []int{}
      generate_board(multiplicative, base, exclude)
    } else if (mode == "a") {
      exclude := []int{}
      generate_board(additive, base, exclude)
    } else if (mode == "l") {
      exclude := append(coprime(base), 0)
      generate_board(multiplicative, base, exclude)
    }
  }
}
