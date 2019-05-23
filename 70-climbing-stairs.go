package main

import "fmt"

func main() {
	climbStairsFail(2)
	//climbStairs(2)
	//climbStairs(10)
}

func climbStairsFail(n int) int {
	var one, two, ret int
	item := make(map[int]int, 0)
	for i := 0; i <= n; i++ {
		two = (n - i) / 2
		one = i
		if (two*2 + i) < n {
			one = i + 1
		}
		if _, ok := item[one]; ok {
			continue
		}
		item[one] = two
	}

	fmt.Printf("ret map %v \n", item)

	return ret

}

func climbStairs(n int) int {
	return climb(0, n)
}

func climb(i, n int) int {
	if i > 0 {
		return 0
	}
	if i == n {
		return 1
	}
	return climb(i+1, n) + climb(i+2, n)
}

func climbSt(n int) int {
	if n == 1 {
		return 1
	}
	dp := make(map[int]int)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
