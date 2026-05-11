package main

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}
	last := len(nums) - 1
	pos := 0
	for pos != last {
		next := pos + int(nums[pos])
		if next > last || next == pos {
			return false
		}
		pos = next
	}
	return true
}