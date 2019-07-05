package main

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	var ans []string

	for num :=1; num <=n;num++ {
		divisibleBy3 := (num % 3 == 0)
		divisibleBy5 := (num % 5 == 0)

		if divisibleBy3 && divisibleBy5 {
			ans = append(ans,"FizzBuzz")
		} else if divisibleBy3 {
			ans = append(ans,"Fizz")
		} else if divisibleBy5 {
			ans = append(ans,"Buzz")
		} else  {
			d := strconv.Itoa(num)
			ans = append(ans,d)
		}
	}
	return ans
}