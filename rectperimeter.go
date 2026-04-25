package main

func RectPerimeter(length, width int) int {
	if length <= 0 || width <= 0 {
		return -1
	}
	perimeter := 2 * (length + width)
	return perimeter
}

// The RectPerimeter function calculates the perimeter of a rectangle given its length and width. 
// It first checks if either the length or width is less than or equal to 0. If so, it returns -1 to indicate an invalid input. 
// Otherwise, it calculates the perimeter using the formula 2 * (length + width) and returns the result.