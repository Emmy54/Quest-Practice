package main 

func FishAndChips(nb int) string {
	if nb < 0 {
		return "Invalid Input"
	}
	if nb%2 == 0 && nb%3 == 0 {
		return "Fish and Chips"
	}
	if nb%2 == 0 {
		return "Fish"
	}
	if nb%3 == 0 {
		return "Chips"
	}
	return "Nothing"
}