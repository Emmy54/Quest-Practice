package main

func LastWord(s string) string {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' { // Skip trailing spaces
		i--
	}
	result := ""
	for i >= 0 && s[i] != ' ' { // Collect characters of the last word
		result = string(s[i]) + result
		i--
	}
	return result
}	