package main

import (
	"fmt"
)

func moveZeroes(nums []int)  {
	var (
		tmpL  = make([]int, 0)
		numZeroes int
	)

	for i:= 0; i< len(nums);i++ {
		if nums[i] == 0 {
			numZeroes++
		} else {
				tmpL = append(tmpL, nums[i])
		}
	}
	fmt.Printf("tmpl%v", tmpL)
	for numZeroes >0 {
		tmpL = append(tmpL, 0)
		numZeroes--
	}
	for i:= 0;i<len(nums);i++ {
		nums[i] = tmpL[i]
	}

}

func main() {
	num := []int{0,1,0,3,12}
	moveZeroes(num)
	fmt.Printf("num: %v",num)
}
