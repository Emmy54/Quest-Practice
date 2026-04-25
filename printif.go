package main

func PrintIf(str string) string {
	if len(str) == 0 || len(str) > 3 {
		return "Valid Input"
	}
	return "Invalid Input"
}

// The PrintIf function checks the length of the input string. 
// If the length is 0 or greater than 3, it returns "Valid Input". 
// Otherwise, it returns "Invalid Input".