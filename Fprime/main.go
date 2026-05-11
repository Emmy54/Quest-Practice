package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printInt(n int) {
	if n >= 10 {
		printInt(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}

func atoi(s string) (int, bool) {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n, true
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, ok := atoi(os.Args[1])
	if !ok || n <= 1 {
		return
	}

	first := true
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			if !first {
				z01.PrintRune('*')
			}
			printInt(i)
			first = false
			n /= i
		}
	}
	z01.PrintRune('\n')
}
