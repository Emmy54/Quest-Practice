package main

func ConcatSlice(slice1 []int, slice2 []int) []int {
	return append(slice1, slice2...)
}