package main

import (
	"fmt"
)

//回文字符串的个数

//dp[i][j]=dp[i+1][j-1] 参考dp.png图片

func countSubstrings(s string) int {
	if s == "" {
		return 0
	}
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	//初始化dp
	result := 0
	//
	fmt.Println(dp)
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
				result++
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
				result++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				result++
			}
		}
	}
	return result
}
