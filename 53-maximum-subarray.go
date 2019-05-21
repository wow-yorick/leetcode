package main

import "fmt"

func main() {
	n := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	r := maxSubArray(n)
	fmt.Printf("the result: %d", r)
}

func maxSubArray(nums []int) int {
	var (
		max, min, ret int
	)
	ret = nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > ret {
			ret = nums[i]
		}
		tmp := 0
		flg := true
		for j := i + 1; j < len(nums); j++ {
			if flg {
				tmp += nums[i]
			}
			flg = false
			tmp += nums[j]
			if tmp > ret {
				ret = tmp
			}
		}
	}
	return ret
}

func maxSubArrayRight(nums []int) int {
	max, sum := nums[0], nums[0]
	for _, v := range nums[1:] {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
