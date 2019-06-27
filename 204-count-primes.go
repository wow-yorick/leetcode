package main

import (
	"fmt"
	"math"
)

func main_15() {
	n := countPrimes(10)
	fmt.Printf("xx:%v", n)
}

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 || n == 4 {
		return 2
	}
	//大于3的数至少有2个质数
	count := 2
	prim := []int{2, 3}
	for i := 5; i <= n; i++ {
		if isPrime(i) {
			count++
			prim = append(prim, i)
		}
	}
	return count

}

func isPrime(num int) bool {
	if num <= 3 {
		return num > 1
	}

	//不在6的倍数两侧的一定不是质数
	if num%6 != 1 && num%6 != 5 {
		return false
	}
	//假如n是合数，必然存在非1的两个约数p1和p2，其中p1<=sqrt(n)，p2>=sqrt(n)
	sqrt := int(math.Sqrt(float64(num)))
	for i := 5; i <= sqrt; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}

	}

	return true
}
