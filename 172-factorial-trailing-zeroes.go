package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 	ret := trailingZeroes(42)
	// 	println(ret)
	str := "Hello,世界"
	//方法一：格式化打印
	for _, ch1 := range str {
		fmt.Printf("%q \n", ch1) //单引号围绕的字符字面值，由go语法安全的转义
	}
	fmt.Println("==========>方法二\n")
	//方法二：转化输出格式
	for _, ch2 := range str {
		fmt.Printf("%s", reflect.TypeOf(ch2))
		fmt.Println(string(ch2))
	}
}

func trailingZeroes(n int) int {
	rst := 0
	i := 5
	for n/i != 0 {
		rst += n / i
		i *= 5
	}
	return rst
}

func trailingZeroesTwo(n int) int {
	if n/5 == 0 {
		return 0
	}
	return n/5 + trailingZeroes(n/5)
}
