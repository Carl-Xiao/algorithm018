package main

//dfs 题目
//上下左右斜线的8个方法 ，这个是需要dfs 判断当前是否需要改变
var dx = []int{-1, 0, 1, -1, 1, 0, 1, -1}
var dy = []int{-1, 1, 1, 0, -1, -1, 0, 1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]
	//遇到炸弹直接返回
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	dfs(board, x, y)
	return board
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'E' {
		return
	}
	num := getNumsOfBombs(board, x, y)
	if num == 0 {
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			dfs(board, nx, ny)
		}
	} else {
		board[x][y] = byte('0' + num)
	}
}

func getNumsOfBombs(board [][]byte, x, y int) int {

	num := 0
	for i := 0; i < 8; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || nx >= len(board) || ny < 0 || ny >= len(board[0]) {
			continue
		}
		if board[nx][ny] == 'M' || board[nx][ny] == 'X' {
			num++
		}
	}
	return num
}

func main() {

}
