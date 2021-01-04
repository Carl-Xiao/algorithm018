package main

//状态转移方程dp代表长度
// dp[i]=max(dp[j]+1,dp[i])

//n平方的时间复杂度
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var res int
	//初始化
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	//处理
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
