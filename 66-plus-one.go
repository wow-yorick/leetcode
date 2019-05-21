package main

import "fmt"

func main() {
	ret := plusOne([]int{1, 2, 3, 9})
	fmt.Printf("ret %v", ret)
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{0}
	}
	if digits[0] == 0 {
		return []int{1}
	}

	ret := []int{}

	ret = append(ret, (digits[len(digits)-1]+1)%10)
	carry := (digits[len(digits)-1] + 1) / 10
	for i := len(digits) - 2; i >= 0; i-- {
		if carry > 0 {
			ret = append(ret, (digits[i]+carry)%10)
			carry = (digits[i] + carry) / 10
		} else {
			ret = append(ret, digits[i])
			carry = 0
		}
	}
	if carry > 0 {
		ret = append(ret, carry)
	}
	for x := 0; x < len(ret); x++ {
		if x < len(ret)-1-x {
			ret[x], ret[len(ret)-1-x] = ret[len(ret)-1-x], ret[x]
		}
	}
	return ret
}
