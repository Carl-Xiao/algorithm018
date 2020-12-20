package main

import "fmt"

// 一条包含字母 A-Z 的消息通过以下方式进行了编码：

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。

// 题目数据保证答案肯定是一个 32 位的整数。
//状态转移方程
//数量 dp
//当前等于0的时候,前缀一定是1或者2

// 处理 最开始比较特别所以一定要注意边界值 i>2=2  dp[i]=dp[i-2]  dp[i]=1
//不等于0 dp[i]=dp[i]+dp[i-1]
func numDecodings(s string) int {
	if s == "0" {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == '1' || s[i-1] == '2' && s[i] >= '0' && s[i] <= '6' {
			if i >= 2 {
				dp[i] = dp[i-2]
			} else {
				dp[i] = 1
			}
		}
		if s[i] != '0' {
			dp[i] += dp[i-1]
		}
	}
	fmt.Println(dp)
	return dp[len(s)-1]
}
func main() {
	numDecodings("12")
}
