package main

import (
	"sort"
)

func intersectFail(nums1 []int, nums2 []int) []int {
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

func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	if len(nums1) != 0 && len(nums2) != 0 {
		for _, a:= range nums1 {
			for i,b := range nums2 {
				if a == b {
					result = append(result, a)
					nums2 = append(nums2[:i], nums2[i+1:]...)
					break
				}
			}
		}

	} else {
		return result
	}

	return result
}