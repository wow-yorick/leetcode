package main

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	rtm := make(map[int]int, 0)
	for _, v := range nums {
		rtm[v] += 1
	}
	for n, t := range rtm {
		if t == 1 {
			return n
		}
	}
	return 0
}

func singleNumberNew(nums []int) int {
	var ret = []int{}

	inArray := func(e int, ls []int) bool {
		for _, l := range ls {
			if e == l {
				return true
			}
		}
		return false
	}
	removeEl := func(e int, ls []int) []int {
		for i, v := range ls {
			if e == v {
				ls = append(ls[:i], ls[i+1:]...)
				return ls
			}
		}
		return ls
	}
	for _, v := range nums {
		if !inArray(v, ret) {
			ret = append(ret, v)
		} else {
			ret = removeEl(v, ret)
		}
	}
	return ret[0]
}
