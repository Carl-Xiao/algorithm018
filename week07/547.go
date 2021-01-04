package main

//bfs 或者并查集
//并查集 constructor(构造函数) find  uniton(合并)

//https://www.bilibili.com/video/BV13t411v7Fs?p=1 讲的很好
type union struct {
	count  int
	parent []int
}

func constructor(n int) *union {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &union{
		parent: parent,
		count:  n,
	}
}

func (u *union) find(p int) int {
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

func (u *union) uniton(p, q int) {
	rootQ := u.find(q)
	rootP := u.find(p)

	if rootP == rootQ {
		return
	}
	u.parent[rootQ] = rootP
	u.count--
}

func findCircleNum(M [][]int) int {
	union := constructor(len(M))
	for i := 0; i < len(M); i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				union.uniton(i, j)
			}
		}
	}
	return union.count
}
