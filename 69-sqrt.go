package main

import "fmt"

func main() {
	x := 4
	y := mySqrt(x)

	fmt.Printf("x:%d y:%d", x, y)

	x = 5
	y = mySqrt(x)

	fmt.Printf("x:%d y:%d", x, y)

	x = 100
	y = mySqrt(x)

	fmt.Printf("x:%d y:%d", x, y)
}

func mySqrtFail(x int) int {
	if x <= 0 {
		return 0
	}
	precision := 1
	ret := x / 2
	for {
		if ret*ret > x {
			ret = ret / 2
		} else {
			if (x - ret*ret) <= precision {
				break
			}
			ret++
		}
	}
	return ret
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	approx := 0.5 * float64(x)
	for i := 0; i < 20; i++ {
		betterapprox := 0.5 * float64(approx+float64(x)/approx)
		approx = betterapprox
	}
	return int(approx)

}
