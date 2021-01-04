package main

//dfs或者bfs 四联通向量处理

//dfs
func numIslands(grid [][]byte) int {
	count := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j, grid)
			}
		}
	}
	return count
}

func dfs(i, j int, grid [][]byte) {
	if i < 0 || j < 0 || i > len(grid) || j > len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(i-1, j, grid)
	dfs(i+1, j, grid)
	dfs(i, j-1, grid)
	dfs(i, j+1, grid)
}

//并查集处理
//直接copy前面的代码部分
type gridUnion struct {
	count  int
	parent []int
}

//二维数组变一维  index=i*列+j
func constructorGrid(grid [][]byte) *gridUnion {
	row := len(grid)
	col := len(grid[0])
	count := row * col
	u := &gridUnion{parent: make([]int, count)}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			u.parent[i*col+j] = i*col + j
			if grid[i][j] == '1' {
				u.count++
			}
		}
	}
	return u
}

func (u *gridUnion) find(p int) int {
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

func (u *gridUnion) uniton(p, q int) {
	rootQ := u.find(q)
	rootP := u.find(p)

	if rootP == rootQ {
		return
	}
	u.parent[rootQ] = rootP
	u.count--
}

func getIndex(i, j, n int) int {
	return i*n + j
}

func numIslandsUnion(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	unionFind := constructorGrid(grid)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				if i-1 >= 0 && grid[i-1][j] == '1' {
					unionFind.uniton(i*col+j, (i-1)*col+j)
				}
				if i+1 < row && grid[i+1][j] == '1' {
					unionFind.uniton(i*col+j, (i+1)*col+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					unionFind.uniton(i*col+j, i*col+(j-1))
				}
				if j+1 < col && grid[i][j+1] == '1' {
					unionFind.uniton(i*col+j, i*col+(j+1))
				}
				//充值为0
				grid[i][j] = '0'
			}
		}
	}
	return unionFind.count
}
