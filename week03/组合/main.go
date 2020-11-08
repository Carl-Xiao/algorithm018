package main

import "fmt"

//NO 77
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
// 心里是蒙蔽的，咋的这递归完全就不一样了，思考方式都不同了
// 看了第三遍的题解 有一点点感觉,但还是理解不了
func combine(n int, k int) [][]int {
	results := [][]int{}
	if k > n {
		return results
	}

	dfs(n, k, 1, []int{}, &results)
	return results
}

func dfs(n int, k, begin int, buf []int, results *[][]int) {
	if k == 0 {
		tmp := make([]int, len(buf))
		copy(tmp, buf)
		*results = append(*results, tmp)
		return
	}
	for i := begin; i <= n; i++ {
		buf = append(buf, i)
		dfs(n, k-1, i+1, buf, results)
		buf = buf[:len(buf)-1]
	}
}

func main() {
	fmt.Println(combine(4, 2))
}
