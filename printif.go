package main

func PrintIf(str string) string {
	if len(str) == 0 || len(str) > 3 {
		return "Valid Input"
	}
	return "Invalid Input"
}