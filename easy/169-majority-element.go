package main

func majorityElement(nums []int) int {
	retMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, ok := retMap[nums[i]]
		if ok {
			retMap[nums[i]] = retMap[nums[i]] + 1
		} else {
			retMap[nums[i]] = 1
		}
	}
	max := 0
	maxV := 0
	for k, v := range retMap {
		if v > max {
			max = v
			maxV = k
		}
	}
	return maxV
}
