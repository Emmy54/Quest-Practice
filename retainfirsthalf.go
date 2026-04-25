package main

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return "Invalid Input"
	}
	if str == " " {
		return "Invalid Input"
	}
	HalfLength := len(str) / 2
	return str[:HalfLength]
}

// The RetainFirstHalf function takes a string as input and returns the first half of the string. 
// It first checks if the input string is empty or contains only a space. If either condition is true, it returns "Invalid Input". 
// Otherwise, it calculates the half length of the string and returns the substring from the beginning to the half length.