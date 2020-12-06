package main

import "fmt"

//被围绕的区域 dfs处理应该更快 这个跟岛屿数量问题应该是一样的意思
func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isEdge(i, j, m, n) && board[i][j] == 'O' {
				dfs(i, j, board)
			}
		}
	}
	fmt.Println(board)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func isEdge(i, j, m, n int) bool {
	return i == 0 || j == 0 || i == m-1 || j == n-1
}

func dfs(i, j int, board [][]byte) {
	if i < 0 || j < 0 || i >= len(board)-1 || j >= len(board[0])-1 || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}
	board[i][j] = '#'
	dfs(i-1, j, board)
	dfs(i+1, j, board)
	dfs(i, j-1, board)
	dfs(i, j+1, board)
}
