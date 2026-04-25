package main

func PrintIfNot(str string) string {
	if len(str) <= 3 {
		return "Invalid Input"
	}
	return "Valid Input"
}

// this is the reverse of PrintIf. 
// It checks if the length of the input string is less than or equal to 3. 
// If it is, it returns "Invalid Input". Otherwise, it returns "Valid Input".