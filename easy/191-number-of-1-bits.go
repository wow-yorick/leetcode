package main

import "fmt"

func main_8() {
	t := hammingWeight(4)
	fmt.Println("result:", t)
}

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count += int(num & 1)
		num >>= 1
	}
	return count
}
