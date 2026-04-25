package main

func RectPerimeter(length, width int) int {
	if length <= 0 || width <= 0 {
		return -1
	}
	perimeter := 2 * (length + width)
	return perimeter
}