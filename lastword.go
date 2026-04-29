package main

func LastWord(s string) string {
	if s == "" {
		return ""
	}
	i := len(s) - 1 
	for i >= 0 && s[i] == ' ' {
		i--
	}
	end := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	if end < 0 {
		return "\n"
	}
	return s[i+1 : end+1] + "\n"
}

//