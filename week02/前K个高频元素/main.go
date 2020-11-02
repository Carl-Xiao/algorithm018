package main

import (
	"container/heap"
	"fmt"
)

// 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

//想到了用堆来解决

//前K个频率高的元素
func topKFrequent(nums []int, k int) []int {
	mc := make(map[int]int)
	for _, value := range nums {
		mc[value]++
	}

	h := &head{}
	heap.Init(h)

	for key, value := range mc {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type head [][2]int

func (h head) Len() int           { return len(h) }
func (h head) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h head) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *head) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *head) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func main() {
	fmt.Print(topKFrequent([]int{1, 2}, 2))
}
