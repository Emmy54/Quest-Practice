package main 

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}
	count := 0
	for {
		n = n / base
		count++
		if n == 0 {
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

// //Write a function DigitLen() that takes two integers as arguments and returns the times the first int can be divided by the second until it reaches zero.

//     The second int must be between 2 and 36. If not, the function returns -1.
//     If the first int is negative, reverse the sign and count the digits.
