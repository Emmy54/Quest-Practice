package main

func HashCode(dec string) string {
	size := len(dec)
	result := ""
	for _, char := range dec {
		hashed := (int(char) + size) % 127

		if hashed < 33 {
			hashed += 33
		}
		result += string(rune(hashed))
	}
	return result
}

//hello