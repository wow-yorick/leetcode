package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[len(nums) -1] != len(nums) {
		return len(nums)
	} else if nums[0] != 0 {
		return 0
	}
	for i:=0; i<len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return -1
}


func main() {
	t:= missingNumber([]int{3,0,1,4})
	fmt.Printf("ret %d", t)
}