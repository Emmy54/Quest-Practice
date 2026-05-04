package main

func Itoa(num int) string {
	if num == 0 {
		return "0\n"
	}
	sign := ""
	result := ""
	if num < 0 {
		sign = "-"
		num = -num
	}
	for ; num > 0; num /= 10 {
		result = string(rune('0' + num%10)) + result
	}
	return sign + result
}
//what is this?
//This function converts an integer to a string representation. 
// It handles both positive and negative integers, and returns the string with a newline character at the end. 
// If the input number is zero, it simply returns "0\n".