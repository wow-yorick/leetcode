package main

import "fmt"

//暴力
func rotate(nums []int, k int) {
	var previous int
	for i := 0; i < k; i++ {
		previous = nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			nums[j], previous = previous, nums[j]
		}
	}
}

func main() {
	il := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(il, 3)
	fmt.Printf("%+v\n", il)
}

//slice
func rotate2(nums []int, k int) {
	numsLen := len(nums)
	if numsLen != 1 {
		clone := nums[numsLen-(k%numsLen):]
		latter := nums[:numsLen-(k%numsLen)]
		clone = append(clone, latter...)
		copy(nums, clone)
	}
}

//使用额外数组
func rotate3(nums []int, k int) {
	a := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		a[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = a[i]
	}
}

func rotate4(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for start := 0; count < len(nums); start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % len(nums)
			nums[next], prev = prev, nums[next]
			current = next
			count++
			if start == current {
				break
			}
		}
	}
}

func rotate5(nums []int, k int) {
	k = k % len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
