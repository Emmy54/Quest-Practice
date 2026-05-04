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

// Write a function IsCapitalized() that takes a string as an argument and returns true if each word in the string begins with either an uppercase letter or a non-alphabetic character.
// If any of the words begin with a lowercase letter return false.
// If the string is empty return false.