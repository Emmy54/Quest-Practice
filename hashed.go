package main

func Hashed(str string) string {
	if str == "" {
		return ""
	}

	size := len(str) // the size variable is assigned the length of the input string, which will be used to calculate the hashed value for each character in the string.
	var result string // the result variable is initialized as an empty string, which will be used to store the final hashed string that will be returned at the end of the function.
	for _, char := range str {
		hashed := (int(char) + size) % 127 // the hashed variable is calculated by taking the ASCII value of the current character (converted to an integer), adding the size of the string to it, and then taking the result modulo 127. 
		// This ensures that the hashed value stays within the range of valid ASCII characters.
		if hashed < 33 { // if the hashed value is less than 33, it adds 33 to it to ensure that the resulting character is a printable ASCII character.
			result += string(rune(33))
		}
		result += string(rune(hashed)) // the hashed value is converted back to a character using the rune type and added to the result string. 
		// This process is repeated for each character in the input string, 
		// resulting in a new string where each character has been transformed based on its original ASCII value and the length of the input string.
	}
	return result
}
