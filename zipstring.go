package main

import "strconv"

func ZipString(s string) string {
	result := ""
	i := 0
	for i < len(s) {
		count := 1
		for i+count < len(s) && s[i+count] == s[i] {
			count++
		}
		result += strconv.Itoa(count) + string(s[i])
		i += count
	}
	return result
}

//without strconv package

// func ZipString(s string) string {
// 	result := ""
// 	i := 0
// 	for i < len(s) {
// 		count := 1
// 		for i+count < len(s) && s[i+count] == s[i] {
// 			count++
// 		}
// 		result += string(rune('0'+count)) + string(s[i])
// 		i += count
// 	}
// 	return result
// }