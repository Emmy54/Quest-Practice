package main

func ThirdTimeIsACharm(str string) string {
	if str == "" {
		return "\n"
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	return result
}