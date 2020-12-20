package main

//爬楼梯问题

//1 2 3 ...可以爬1 2 3阶
//只能从m[] 数组给的选择
//爬楼梯的中间不能相等

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 注意：给定 n 是一个正整数。
func climbStairs(n int) int {
	var dp []int
	dp[0] = 1
	dp[1] = 2
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//可以爬 1 2 3
//dp[i] = dp[i-1] + dp[i-2]+dp[i-3]

//需要数组的情况下
