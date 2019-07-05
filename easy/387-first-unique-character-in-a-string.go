package main

import (
	"strings"
)

func firstUniqChar(s string) int {
	s = strings.ToLower(s)
	l := len(s)
	countm := make(map[uint8]int,0)
	for i:=0;i<l;i++ {
		countm[s[i]] = countm[s[i]] +1
	}

	for i:=0; i<l;i++ {
		if v,ok := countm[s[i]];ok && v == 1 {
			return i
		}
	}
	return -1
}