package main

func maxProfit2Wrong(prices []int) int {
	var ret int
	buyPrice := int(^uint(0) >> 1)

	minprice := int(^uint(0) >> 1)
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > profit {
			profit = prices[i] - minprice
		}
	}

	for i := 0; i < len(prices); i++ {

		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else if prices[i]-buyPrice > 0 {
			ret += prices[i] - buyPrice
			buyPrice = int(^uint(0) >> 1)
		}
	}

	if profit > ret {
		return profit
	} else {
		return ret
	}

}

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	i := 0
	valley := prices[0]
	peak := prices[0]
	maxprofit := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		maxprofit += peak - valley
	}
	return maxprofit
}
