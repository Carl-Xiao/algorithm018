package main

import "sort"

//因为给定的每个小区间都是升序的，问题就变得很清楚了，首先排序元素，
//然后一次遍历区间的最大值关系，然后决定是否合并
func merge(intervals [][]int) [][]int {
	//排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 没有一点重合
			res = append(res, prev)
			prev = cur
		} else { // 有重合
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
