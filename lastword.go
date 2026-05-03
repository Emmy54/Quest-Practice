package main

func LastWord(s string) string {
	if s == "" {
		return ""
	}
	i := len(s) - 1 
	for i >= 0 && s[i] == ' ' { // find the index of the last space character in the string
		i--
	}
	end := i
	for i >= 0 && s[i] != ' ' { // find the index of the last word 
		i--
	}
	if end < 0 {
		return ""
	}
	return s[i+1 : end+1]
}

//