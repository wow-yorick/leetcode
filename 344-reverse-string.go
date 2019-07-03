package main

import (
	"fmt"
)

func reverseString(s []byte)  {
	for i,j:=0,len(s)-1; i < j; i,j = i+1, j-1 {
		s[i],s[j] = s[j],s[i]
	}
}

func main() {
	st := []byte("hello")
	reverseString(st)
	fmt.Printf("%s", st)
}