package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	j := 0
	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s2[i] == s1[j] {
			j++
		}
	}

	if j == len(s1) {
		for i := 0; i < len(s1); i++ {
			z01.PrintRune(rune(s1[i]))
		}
		z01.PrintRune('\n')
	}
}