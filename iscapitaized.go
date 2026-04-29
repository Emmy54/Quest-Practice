package main 

func IsCapitalized(str string) bool {
	if str == "" {
		return false
	}
	for i, char := range str {
		if (char >= 'a' && char <= 'z') && (i == 0 || (str[i-1] == ' ')) {
			return false
		}
	}
	return true
}