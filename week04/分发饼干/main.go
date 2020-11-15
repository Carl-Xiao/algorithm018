package main

import (
	"fmt"
	"sort"
)

// 455. 分发饼干

//先排序 然后两个数组依次配对
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	fmt.Print(g)
	fmt.Print(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}
