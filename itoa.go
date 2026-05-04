package main

func Itoa(num int) string {
	if num == 0 {
		return "0\n"
	}
	result := ""
	if num < 0 {
		result += "-"
		num = -num
	}
	for ; num > 0; num /= 10 {
		result = string(rune('0' + num%10)) + result // convert the last digit to a character and prepend it to the result string
	}
	return result
}
//what is this?
//This function converts an integer to a string representation. 
// It handles both positive and negative integers, and returns the string with a newline character at the end. 
// If the input number is zero, it simply returns "0\n".