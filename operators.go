package main

type operator func(int, int, int) int

func additive(num1 int, num2 int, base int) int {
	return (mod(mod(num1, base)+mod(num2, base), base))
}

func multiplicative(num1 int, num2 int, base int) int {
	return (mod(mod(num1, base)*mod(num2, base), base))
}
