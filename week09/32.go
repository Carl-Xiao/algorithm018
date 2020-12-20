package main

// 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

//状态转移方程 dp代表数量

//s[i]等于右括号的时候  ....) 这时候只需要找前面一个是否属于左括号 或者前面的前面是否已经有左括号
//紧挨着的左括号  dp[i]=dp[i-1]+2
//挨着的不是左括号,那就继续往前找

//使用栈是最合适的求解，但是在这里顺便练习一下动态规划
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	maxAns := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i]-1 == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}
