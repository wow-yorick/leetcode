package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	ret := twoSum(arr, 9)
	fmt.Printf("the result is: %v", ret)

	arr2 := []int{5, 3, 10, 15, 40}
	ret2 := twoSum(arr2, 25)
	fmt.Printf("the result2 is: %v", ret2)
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
