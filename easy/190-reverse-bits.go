package main

import (
	"fmt"
	"log"
	"strconv"
)

//十进制转二进制字符串 然后进行翻转
func reverseBits(num uint32) uint32 {
	r := fmt.Sprintf("%032b", num)
	nr := []byte(r)
	start := 0
	end := len(nr) - 1
	for start < end {
		nr[start], nr[end] = nr[end], nr[start]
		start++
		end--
	}
	re, err := strconv.ParseInt(string(nr), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return uint32(re)
}

func main_7() {
	var i = uint32(43261596)
	d := reverseBits3(i)
	println("before:", i, "after:", d)
}

//00000010100101000001111010011100
func reverseBits2(num uint32) uint32 {
	rnum := uint32(0)
	for i := 0; i < 8; i++ {
		rnum = rnum<<1 + (num & 1)
		fmt.Println(rnum)
		rnum = rnum<<1 + (num >> 1 & 1)
		fmt.Println(rnum)
		rnum = rnum<<1 + (num >> 2 & 1)
		fmt.Println(rnum)
		rnum = rnum<<1 + (num >> 3 & 1)
		fmt.Println(rnum)
		num >>= 4
	}
	return rnum
}

//按位翻转
func reverseBits3(num uint32) uint32 {
	var ret = uint32(0)
	for i := 0; i < 32; i++ {
		ret = ret<<1 + (num & 1)
		// fmt.Println(ret)
		num = num >> 1
		// fmt.Printf("%032b\n", num)
	}
	return ret
}
