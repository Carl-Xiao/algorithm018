package main

import "fmt"

//转移方程

//dp[i][0] 没有  max(dp[i-1][0],dp[i-1][1]+prices[i]) 不动 卖出
//dp[i][1] 持有  max(dp[i-1][1],dp[i-1][0]-prices[i]) 不动 买进
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = 0 - prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	fmt.Println(dp)
	return dp[len(prices)-1][0]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
