package main

func CountChar(s string) int { // Count the number of characters in a string, excluding spaces
	count := 0 // Initialize a counter to keep track of the number of characters
	for _, char := range s { // Iterate through each character in the string
		if char != ' ' {	// Check if the character is not a space
			count++ // If it's not a space, increment the counter
		}
	}
	return count	// Return the final count of characters
}		