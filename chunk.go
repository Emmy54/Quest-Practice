package main

import "github.com/01-edu/z01"

func printInt(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printInt(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}

func Chunk(slice []int, size int) {
	if len(slice) == 0 || size <= 0 {
		return
	}

	z01.PrintRune('[')
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		z01.PrintRune('[')
		for j := i; j < end; j++ {
			printInt(slice[j])
			if j < end-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(']')
		if i+size < len(slice) {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}