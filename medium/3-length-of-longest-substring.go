package medium

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	var (
		max_length  = 0 //滑动窗口数组
		left, right = 0, 0
	)
	utf8Ind := func(s string, c rune) int {
		sr := []rune(s)
		for i, oc := range sr {
			if c == oc {
				return i
			}
		}
		return -1
	}
	str2 := []rune(s)
	for _, c := range str2 {
		//fmt.Printf("%s contains %c result:%t\n", string(str2[left:right]), c,strings.ContainsAny(string(str2[left:right]), string(c)))
		if !strings.ContainsAny(string(str2[left:right]), string(c)) {
			right += 1
		} else {
			left += utf8Ind(string(str2[left:right]), c) + 1
			//fmt.Printf("left position :%d\n" ,left)
			right += 1
		}
		if right-left > max_length {
			max_length = right - left
		}
	}
	if max_length > 0 {
		return max_length
	} else {
		return len(str2)
	}

}
