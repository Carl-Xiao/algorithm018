package main

import "fmt"

//dp 动态规划
//状态转移方程 dp[i]=dp[i-1]+dp[i-2]
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
func main() {
	fmt.Println(climbStairs(3))
}
