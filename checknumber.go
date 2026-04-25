package main

func CheckNumber(str string) bool {
	for _, char := range str {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}