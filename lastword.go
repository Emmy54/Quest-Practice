package main

// func LastWord(s string) string {
// 	if s == "" {
// 		return ""
// 	}
// 	i := len(s) - 1 
// 	for i >= 0 && s[i] == ' ' { // find the position of the last non-space character
// 		i--
// 	}
// 	end := i
// 	for i >= 0 && s[i] != ' ' { // find the position of the last space character before the last word
// 		i--
// 	}
// 	if end < 0 {
// 		return ""
// 	}
// 	return s[i+1 : end+1]
// }

//
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