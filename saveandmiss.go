package main

func SaveAndMiss(s string, n int) string {
	if n <= 0 {
		return s
	}

	result := ""
	save := true

	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		if save {
			result += s[i:end]
		}
		save = !save
	}

	return result
}