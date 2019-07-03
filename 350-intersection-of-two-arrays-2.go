package main

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var use []int
	var comp []int
	if len(nums1) < len(nums2) {
		use = nums1
		comp = nums2
	} else {
		use = nums2
		comp = nums1
	}
	inArray := func(n int, nums []int) bool {
		for _,v :=range nums {
			if n == v {
				return true
			}
		}
		return false
	}

	var ret []int
	for i:=0;i<len(use) ;i++ {
		if inArray(use[i], comp) {
			ret = append(ret, use[i])
		}
	}
	return ret

}