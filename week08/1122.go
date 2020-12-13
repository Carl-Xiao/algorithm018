package main

import "sort"

//日常读不懂题目！！！
// arr2出现 但是arr1未出现元素,升序存放到元素后面
//新开一块空间存放数据
// 输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
// 输出：[2,2,2,1,4,3,3,9,6,7,19]
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int][]int)
	for _, v := range arr2 {
		m[v] = make([]int, 0)
	}
	var other []int
	for _, v := range arr1 {
		if l, ok := m[v]; ok {
			l = append(l, v)
			m[v] = l
		} else {
			other = append(other, v)
		}
	}
	// 将不在arr2的元素存在另外一个数组中，对这个数组重新排序
	sort.Ints(other)
	var res []int
	for _, v := range arr2 {
		res = append(res, m[v]...)
	}
	//添加到arr2中
	res = append(res, other...)

	return arr1
}
