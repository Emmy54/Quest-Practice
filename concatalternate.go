package main

func ConcatAlternate(slice1 []int, slice2 []int) []int {
	result := []int{}

	minLen := len(slice1)
	if len(slice2) < minLen {
		minLen = len(slice2)
	}

	for i := 0; i < minLen; i++ {
		if len(slice1) >= len(slice2) {
			result = append(result, slice1[i], slice2[i])
		} else {
			result = append(result, slice2[i], slice1[i])
		}
	}

	// append remainder of the longer slice
	result = append(result, slice1[minLen:]...)
	result = append(result, slice2[minLen:]...)

	return result
}