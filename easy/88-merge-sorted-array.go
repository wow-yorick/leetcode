package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

func mergeOther(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p--
		p2--
	}
}

func mergeOther2(nums1 []int, m int, nums2 []int, n int) {
	var tmp int
	for n > 0 {
		if m == 0 || nums2[n-1] > nums1[m-1] {
			n = n - 1
			tmp = nums2[n]
		} else {
			m = m - 1
			tmp = nums1[m]
		}

		nums1[m+n] = tmp
	}
}
