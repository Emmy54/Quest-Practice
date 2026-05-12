package main 

// func FirstWord(str string) string {
// 	if str == "" {
// 		return str
// 	}
// 	firstWord := ""
// 	for _, char := range str {
// 		if char == ' ' {
// 			break
// 		}
// 		firstWord += string(char)
// 	}
// 	return firstWord
// }

//this function takes a string as input and returns the first word in the string.
// it first checks if the input string is empty, if it is, it returns an empty string.
// then it initializes an empty string variable called firstWord to store the characters of the first word.
// it uses a loop to iterate through each character in the input string, if it encounters a space character, it breaks the loop, otherwise it adds the character to the firstWord variable.
// finally, it returns the firstWord variable which contains the first word from the input string.


// the above implementation does not handle leading spaces correctly, it will return an empty string if the first character is a space.
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

// func FirstWord (str string) string {
// 	w := ""
// 	for _, c := range str {
// 		if c == ' ' && w != "" {
// 			break
// 		} else if c != ' ' {
// 			w += string(c)
// 		}
// 	}
// 	return w 
// }