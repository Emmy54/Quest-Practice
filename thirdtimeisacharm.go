package main

func ThirdTimeIsACharm(str string) string {
	if str == "" {
		return "\n"
	}

	result := ""
	rune := []rune(str)
	for i := 2; i < len(rune); i += 3 {
		result += string(rune[i])
	}
	return result + "\n"
}