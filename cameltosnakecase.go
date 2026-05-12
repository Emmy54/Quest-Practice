package main 

func CamelToSnakeCase(str string) string {
	if !isCamelCase(str) {
		return str
	}
	snakeCase := ""
	for i, char := range str {
		if char >= 'A' && char <= 'Z' {
			if i != 0 {
				snakeCase += "_"
			}
			snakeCase += string(char)	//
		} else {
			snakeCase += string(char)
		}
	}
	return snakeCase
}
// this camelcase helper function checks if the input string is in camel case format.
func isCamelCase(str string) bool {
	if str == "" {
		return false
	}
	for i, char := range str {
		if char >= '0' && char <= '9' {
			return false
		}
		if char == ' ' || (char < 'a' && char > 'z') || (char < 'A' && char > 'Z') {
			return false
		}
		if char >= 'A' && char <= 'Z' {
			if i == len(str)-1 { //checks the last character of the string to ensure it is not uppercase
				return false
			}
			if i > 0 {
				prevChar := str[i-1]  //
				if prevChar >= 'A' && prevChar <= 'Z' {
					return false
				}
			}
		}
	}
	return true
}

// the CamelToSnakeCase function takes a string as input and converts it from camel case to snake case format.
// i created an helper function called isCamelCase to check if the input string is in camel case format, if it is not, it returns "Invalid Input".
// then i initialize an empty string variable called snakeCase to store the characters of the snake case result.
// i use a loop to iterate through each character in the input string, if it encounters an uppercase character, it adds an underscore before it (except for the first character) and converts it to lowercase, otherwise it adds the character as is to the snakeCase variable.
// finally, it returns the snakeCase variable which contains the converted string in snake case format.