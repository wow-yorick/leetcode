package main

import "fmt"

func main() {
	retDay := maxProfit([]int{7, 1, 5, 3, 6, 4})
	//retDay := maxProfit([]int{7, 6, 4, 3, 1})
	//retDay := maxProfit([]int{1, 2})
	fmt.Printf("ret day is %d", retDay)
}

func maxProfitWrong(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var buy, buyD, sell, sellD int
	for i := 0; i < len(prices)-1; i++ {
		if i == 0 {
			buy = prices[i]
			buyD = i
		}
		if prices[i] < buy && i < sellD {
			buy = prices[i]
			buyD = i
		}
		for j := i + 1; j <= len(prices)-1; j++ {
			if sell-buy < prices[j]-buy {
				fmt.Printf("profit %d the day %d \n", prices[j]-buy, j+1)
				sell = prices[j]
				sellD = j
			}
		}
	}

	fmt.Printf("sell - buy: %d  sell: %d sellD: %d buy: %d buyD: %d \n", sell-buy, sell, sellD, buy, buyD)
	if sell-buy < 0 || buyD > sellD {
		return 0
	}

	return sell - buy
}

func maxProfit(prices []int) int {
	minprice := int(^uint(0) >> 1)
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > profit {
			profit = prices[i] - minprice
		}
	}
	return profit
}
