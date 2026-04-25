package main

func GCD(a, b int) int {
	if a == 0 || b == 0 {
		return -1
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}