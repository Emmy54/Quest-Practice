package main 

// the former implementation does not handle leading spaces correctly, it will return an empty string if the first character is a space.
// the following implementation handles leading spaces by first skipping them before extracting the first word.
func FirstWord(str string) string {
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	result := ""
	for i < len(str) && str[i] != ' ' {
		result += string(str[i])
		i++
	}
	return result
}
