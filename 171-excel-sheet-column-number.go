package main

import (
	"fmt"
	"strings"
	"unicode"
)

func titleToNumberLimit(s string) int {
	if len(s) < 1 {
		return -1
	}
	upStr := strings.ToUpper(s)
	strToIndMap := func() map[rune]int {
		ret := make(map[rune]int)
		for i := 0; i < 26; i++ {
			t := 'A' + i
			ret[rune(t)] = i + 1
		}
		return ret
	}
	stm := strToIndMap()
	if _, ok := stm[rune(upStr[0])]; !ok {
		return -1
	}
	if len(s) == 1 {
		return stm[rune(upStr[0])]
	} else if len(s) == 2 {
		if !unicode.IsLetter(rune(s[1])) {
			return -1
		}
		if _, ok := stm[rune(upStr[1])]; !ok {
			return -1
		}

		return stm[rune(upStr[0])]*26 + stm[rune(upStr[1])]

	} else if len(s) == 3 {
		if !unicode.IsLetter(rune(s[1])) {
			return -1
		}
		if _, ok := stm[rune(upStr[1])]; !ok {
			return -1
		}
		if !unicode.IsLetter(rune(s[2])) {
			return -1
		}
		if _, ok := stm[rune(upStr[2])]; !ok {
			return -1
		}

		return stm[rune(upStr[0])]*676 + stm[rune(upStr[1])]*26 + stm[rune(upStr[2])]

	}

	//fmt.Printf("string 1 %T, strtomap: %T", upStr[0], strToIndMap()['A'])

	return -1
}

func main() {
	r := titleToNumber("BB")
	fmt.Printf("aa index is %d", r)
}

func titleToNumber(s string) int {
	result := int(0)
	for i := 0; i < len(s); i++ {
		result *= 26
		fmt.Printf("res: %d \n", result)
		result += int(s[i]) - int('A') + 1
	}
	return int(result)
}
