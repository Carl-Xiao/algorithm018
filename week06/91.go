package main

//寻找子问题,处理。证明是证明不了的
// 当s[i] == '0', 若s[i-1]=='1'||'2', 则 dp[i] = dp[i-2]

// 当s[i-1] == '1', dp[i] = dp[i-1] + dp[i-2] //相当于 跨一步 + 跨两步

//最大字母为为26
// 当s[i-1] == '2', 若s[i]<='6', dp[i] = dp[i-1] + dp[i-2], 否则dp[i]=dp[i-1]

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			}
			cur = pre
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			cur += pre
		}
		pre = tmp
	}
	return cur
}
