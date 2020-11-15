package main

//74. 搜索二维矩阵

//感觉这题与二分也没什么关系
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	colum := len(matrix[0]) - 1
	for row := 0; row < len(matrix); row++ {
		if matrix[row][colum] >= target {
			for i := 0; i <= colum; i++ {
				if matrix[row][i] == target {
					return true
				}
			}
			return false
		}
	}
	return false
}
