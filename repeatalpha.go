package main

func RepeatAlpha(str string) string {
	if str == "" {
		return ""
	}
	result := ""
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			repeatcount := int(char - 'A' + 1) // 'A' is 65 in ASCII, so we subtract 64 to get the correct repeat count
			for i := 0; i < repeatcount; i++ {
				result += string(char)
			}
		} else if char >= 'a' && char <= 'z' {
			repeatcount := int(char - 'a' + 1) // 'a' is 97 in ASCII, so we subtract 96 to get the correct repeat count
			for i := 0; i < repeatcount; i++ {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	return result
}