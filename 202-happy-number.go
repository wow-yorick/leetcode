package main

import (
	"fmt"
)

func isHappy(n int) bool {
	numSl := make([]int, 0)
	for n > 0 {
		t := n % 10
		numSl = append(numSl, t)
		n = n / 10
	}
	if len(numSl) == 1 {
		if numSl[0] == 1 {
			return true
		} else if numSl[0] == 4 {
			return false
		}
	}

	var st int
	for _, v := range numSl {
		st += v * v
	}
	return isHappy(st)
}

func main() {
	ret := isHappy2(1)
	fmt.Printf("isHappy num: %b", ret)
}

func isHappy2(n int) bool {
	var temp int

	if n <= 0 {
		return false
	}
	for n > 0 {
		temp += (n % 10) * (n % 10)
		n = n / 10
		if n == 0 {
			if temp == 1 {
				return true
			}
			if temp == 4 {
				return false
			}
			n = temp
			temp = 0
		}
	}
	return false
}
