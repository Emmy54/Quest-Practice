package main 

func CamelToSnakeCase(str string) string {
	if !isCamelCase(str) {
		return "Invalid Input"
	}
	snakeCase := ""
	for i, char := range str {
		if char >= 'A' && char <= 'Z' {
			if i != 0 {
				snakeCase += "_"
			}
			snakeCase += string(char)
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
		if char >= 0 && char <= 9 {
			return false
		}
		if char >= 'A' && char <= 'Z' {
			if i == len(str)-1 {
				return false
			}
			if i > 0 {
				prevChar := str[i-1]
				if prevChar >= 'A' && prevChar <= 'Z' {
					return false
				}
			}
		}
	}
	return true
}