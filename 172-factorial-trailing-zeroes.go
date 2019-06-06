package main

func main() {
	ret := trailingZeroes(42)
	println(ret)
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
