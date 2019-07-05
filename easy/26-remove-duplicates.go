package main

func removeDuplicates(nums []int) int {
	var (
		i = 0
		j = 1
	)
	for ; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
