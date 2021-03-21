package main

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func coprime(num int) []int {
	coprimes := []int{}
	if num > 2 {
		for i := 2; i < num; i++ {
			if hcf(num, i) != 1 {
				coprimes = append(coprimes, i)
			}
		}
	}

	return coprimes
}

func hcf(first_tmp int, second_tmp int) int {
	var hcfnum int
	for i := 1; i <= first_tmp && i <= second_tmp; i++ {
		if first_tmp%i == 0 && second_tmp%i == 0 {
			hcfnum = i
		}
	}
	return hcfnum
}