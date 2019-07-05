package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main_2() {
	isPalindrome("race a car")
	//fmt.Printf("is eq %v", 'e' == 'a')
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	fmt.Printf("the s: %s \n", s)
	i := 0
	j := len(s) - 1
	for i < j {
		if !(unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i]))) {
			i++
			continue
		}
		if !(unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j]))) {
			j--
			continue
		}

		fmt.Printf("sj %v si %v \n", s[j], s[i])
		if s[j] == s[i] {
			i++
			j--
			continue
		}
		return false

	}
	return true
}
