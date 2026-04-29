package main 

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if nb == 2 {
		return nb
	}
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i <= num; i += 6 { // checks for factors from 5 to the square root of num, skipping even numbers
		if num%i == 0 || num%(i+2) == 0 { // checks for factors of the form 6k ± 1
			return false
		}
	}
	return true
}