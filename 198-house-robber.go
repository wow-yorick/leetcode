package main

import (
	"math"
)

func main_10() {
}

func rob(nums []int) int {
	var (
		fiTrue  float64
		fiFalse float64
		temp    float64
	)
	for i := 0; i < len(nums); i++ {
		temp = math.Max(fiTrue, fiFalse)
		fiTrue = fiFalse + float64(nums[i])
		fiFalse = temp
	}
	rt := math.Max(fiTrue, fiFalse)
	return int(rt)
}

// public int rob(int[] nums) {
//     // F[i, boolean robOrNot]
//     // F[i, true] = F[i-1, false] + money[i]
//     // F[i, flase] = max(F[i-1, true], F[i-1, false])
//     // base:
//     // F[0, true] = money[0];
//     // F[0, false] = 0;
//     int fITrue = 0, fIFalse = 0, temp;
//     for (int i = 0; i < nums.length; i++) {
//         temp = Math.max(fITrue, fIFalse);
//         fITrue = fIFalse + nums[i];
//         fIFalse = temp;
//     }
//     return Math.max(fITrue, fIFalse);
// }
