package main

func numDistinct(s string, t string) int {
	dp := [1001][1001]int{}
	sl, tl := len(s), len(t)
	dp[tl][sl] = 1
	for ti := tl; ti >= 0; ti-- {
		for si := sl - 1; si >= 0; si-- {
			dp[ti][si] = dp[ti][si+1]
			if ti == tl {
				dp[ti][si] = 1
			} else if t[ti] == s[si] {
				dp[ti][si] += dp[ti+1][si+1]
			}
		}
	}
	return dp[0][0]
}
