package main

import (
	"os"

	"github.com/01-edu/z01"
)

func contains(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	var seen [256]bool

	for i := 0; i < len(s1); i++ {
		c := s1[i]
		if contains(s2, c) && !seen[c] {
			z01.PrintRune(rune(c))
			seen[c] = true
		}
	}
	z01.PrintRune('\n')
}