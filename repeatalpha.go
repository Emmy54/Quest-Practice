package main

func RepeatAlpha(str string) string {
	if str == "" {
		return ""
	}
	result := ""
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			repeatcount := int(char - 'A' + 1)
			for i := 0; i < repeatcount; i++ {
				result += string(char)
			}
		} else if char >= 'a' && char <= 'z' {
			repeatcount := int(char - 'a' + 1)
			for i := 0; i < repeatcount; i++ {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	return result
}