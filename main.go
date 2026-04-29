package main

import (
	"fmt"
)
func main() {
	fmt.Println(CountAlpha("Hello123"))
	fmt.Println(CountAlpha("12345"))
	fmt.Println(CheckNumber("Hello123"))
	fmt.Println(PrintIf("Bankai"))
	fmt.Println(PrintIfNot("bat"))
	fmt.Println(RectPerimeter(5, 3))
	fmt.Println(RetainFirstHalf("bankai"))
	fmt.Println(CamelToSnakeCase("IsThereMoreLikeUs"))
	fmt.Println(GCD(48, 18))
	fmt.Println(DigitLen(255, 16))
	fmt.Println(FirstWord("Hello World"))
	fmt.Println(FishAndChips(6))
	fmt.Println(Hashed("bankai"))
	fmt.Println(RepeatAlpha("bankai"))
	fmt.Println(FindPrevPrime(12))
	fmt.Println(FromTo(1, 10))
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(LastWord("Hello World   "))
	fmt.Println(Itoa(37))
	fmt.Println(WeAreUnique("We will rock it", "We are the champions, my friend"))
}