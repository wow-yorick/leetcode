package main

func reverse(x int) int {
	var nums, newnums int
	for x != 0 {
		a := x % 10
		newnums = nums*10 + a

		nums = newnums
		x = x / 10

		maxInt32 := 1<<31 - 1
		minInt32 := -1 << 31
		if nums > maxInt32 || nums < minInt32 {
			return 0
		}
	}
	return nums
}
