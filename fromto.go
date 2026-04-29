package main 

import "strconv"

func FromTo(from, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid Input\n"
	}
	result := ""
	start := from
	end := to
	step := 1

	if from > to {
		step = -1
	}

	if from == to {
		if from < 10 {
			result += "0" + strconv.Itoa(from) + "\n"
		} else {
			result += strconv.Itoa(from) + "\n"
		}
	}

	for i := start; ; i += step {
		if i < 10 {
			result += "0" + strconv.Itoa(i)
		} else {
			result += strconv.Itoa(i)
		}
		if i == end {
			break
		}
		result += ", "
	}
	// result += "\n"
	return result
}