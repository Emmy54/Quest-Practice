package main

import "github.com/01-edu/z01"

func PrintRevComb() {
	for a := '9'; a >= '0'; a-- {
		for b := a-1; b >= '0'; b-- {
			for c := b-1; c >= '0'; c-- {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				

				if a != '2' || b != '1' || c != '0' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}