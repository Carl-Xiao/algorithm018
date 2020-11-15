package main

// 122. 买卖股票的最佳时机 II

//这个题就是期望今天一直比昨天的行情好
func maxProfit(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
