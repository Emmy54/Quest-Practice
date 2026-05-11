package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	var seen [256]bool

	for _, s := range os.Args[1:] {
		for i := 0; i < len(s); i++ {
			c := s[i]
			if !seen[c] {
				z01.PrintRune(rune(c))
				seen[c] = true
			}
		}
	}
	z01.PrintRune('\n')
}