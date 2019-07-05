package medium

func searchRange(nums []int, target int) []int {
	targetRange := []int{-1,-1}
	leftIdx := extremeInsertionIndex(nums, target, true)

	if leftIdx == len(nums) || nums[leftIdx] != target {
		return targetRange
	}
	targetRange[0] = leftIdx
	targetRange[1] = extremeInsertionIndex(nums,target, false) -1
	return targetRange

}

func extremeInsertionIndex(nums []int, target int, left bool) int {
	var (
		lo = 0
		hi = len(nums)
	)

	for lo < hi {
		mid := (lo + hi)/2
		if nums[mid] > target || (left &&  target == nums[mid]) {
			hi = mid
		} else {
			lo = mid +1
		}
	}

}