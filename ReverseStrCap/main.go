package main

import (
	"os"

	"github.com/01-edu/z01"
)

func toUpper(c byte) rune {
	if c >= 'a' && c <= 'z' {
		return rune(c - 32)
	}
	return rune(c)
}

func toLower(c byte) rune {
	if c >= 'A' && c <= 'Z' {
		return rune(c + 32)
	}
	return rune(c)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	for i, arg := range args {
		for j := 0; j < len(arg); j++ {
			c := arg[j]
			if c == ' ' {
				z01.PrintRune(' ')
			} else if j == len(arg)-1 || arg[j+1] == ' ' {
				z01.PrintRune(toUpper(c))
			} else {
				z01.PrintRune(toLower(c))
			}
		}
		if i < len(args)-1 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
}