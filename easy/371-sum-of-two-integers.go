package main

func getSum(a int, b int) int {

	for b != 0 {
		carry := a & b

		a = a ^ b

		b = int(carry) << 1

	}
	return a
}

func getSum2(a int, b int) int {

	for b != 0 {
		previous, next := a^b, (a&b)<<1
		a, b = previous, next
	}
	return a
}
