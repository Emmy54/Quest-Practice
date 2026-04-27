package main 

func DigitLen(nb, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if nb < 0 {
		nb = -nb
	}
	if nb == 0 {
		return 1
	}
	count := 0
	for {
		count++
		nb = nb / base
		if nb == 0 {
			break
		}
	}
	return count
}

//first i check if the base is valid, if not i return -1. 
// then i check if the number is negative, if it is i convert it to positive.
// if the number is 0, i return 1 because 0 has one digit in any base.
// then i initialize a count variable to keep track of the number of digits.
// i use a loop to divide the number by the base until it becomes 0, incrementing the count each time.
// finally, i return the count which represents the number of digits in the given base.