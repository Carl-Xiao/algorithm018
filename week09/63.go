package main

//dp[i][j]=dp[i-1][j]+dp[i][j-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if m > 0 && n > 0 && obstacleGrid[0][0] == 1 {
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				obstacleGrid[i][j] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				if i == 0 {
					obstacleGrid[i][j] = obstacleGrid[i][j-1]
				} else if j == 0 {
					obstacleGrid[i][j] = obstacleGrid[i-1][j]
				} else {
					obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
				}
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}
