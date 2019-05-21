package main

import (
	"bytes"
	"strconv"
)

func countAndSay(n int) string {
	if n < 1 || n > 30 {
		return ""
	}
	if n == 1 {
		return "1"
	}

	str := countAndSay(n - 1)

	ret := bytes.NewBuffer([]byte{})
	time := 1
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && str[i] == str[i+1] {
			time++
		} else {
			ret.WriteString(strconv.Itoa(time))
			ret.WriteByte(str[i])

			time = 1
		}
	}
	return ret.String()
}
