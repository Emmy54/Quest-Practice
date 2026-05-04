package main

import "fmt"

func PrintMemory(arr [10]byte) {
	// Print hex values
	for i, char := range arr {
		fmt.Printf("%02x ", char)

		if i < 9 {
			fmt.Print(" ")
		}
		if i == 3 || i == 7 {
			fmt.Println()
		}
	}
	fmt.Println()
	// Print ASCII characters
	for _, char := range arr {
		if char >= 32 && char <= 126 {
			fmt.Printf("%c", char)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}