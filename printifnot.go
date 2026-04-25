package main

func PrintIfNot(str string) string {
	if len(str) <= 3 {
		return "Invalid Input"
	}
	return "Valid Input"
}