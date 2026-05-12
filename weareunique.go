package main

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	instr1 := make(map[rune]bool) // create a map to store the unique characters of str1
	instr2 := make(map[rune]bool)

	for _, char := range str1 {
		instr1[char] = true
	}
	for _, char := range str2 {
		instr2[char] = true
	}

	count := 0
	for char := range instr1 {
		if !instr2[char] {
			count++
		}
	}
	for char := range instr2 {
		if !instr1[char] {
			count++
		}
	}
	return count
}