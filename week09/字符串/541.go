package main

import (
	"fmt"
	"strconv"
)

// 输入:
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]

// 输出: 5

// 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
// 	 返回它的长度 5。

//bfs
func ladderLength(beginWord string, endWord string, wordList []string) int {
	mac := map[string]int{}
	for _, value := range wordList {
		mac[value] = 1
	}
	if mac[endWord] == 0 {
		return 0
	}
	var front []string
	var end []string
	front = append(front, beginWord)
	end = append(end, endWord)
	count := 0
	for len(front) > 0 && len(end) > 0 {
		count++
		if len(front) > len(end) {
			front, end = end, front
		}
		for _, word := range front {
			front = front[1:]
			for index := range word {
				for c := 'a'; c <= 'z'; c++ {
					newWord := word[:index] + string(c) + word[index+1:]
					if contains(newWord, end) {
						fmt.Println(newWord, end)
						return count + 1
					}
					if mac[newWord] == 0 {
						continue
					}
					delete(mac, newWord)
					front = append(front, newWord)
				}
			}
		}
	}
	return 0
}

func contains(key string, result []string) bool {
	for _, value := range result {
		if value == key {
			return true
		}
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := new([]int)
	if root == nil {
		return *result
	}
	dfsVies(root, result, 0)
	return *result
}

func dfsVies(root *TreeNode, result *[]int, depth int) {
	if root == nil {
		return
	}
	if depth == len(*result) {
		*result = append(*result, root.Val)
	}
	depth++
	//注意当前顺序问题====》为解决退化成链表的形式
	dfsVies(root.Right, result, depth)

	dfsVies(root.Left, result, depth)
}

var (
	width map[int]int
	ans   int
)

func widthOfBinaryTree(root *TreeNode) int {
	ans = 0
	width = map[int]int{}
	dfsWidth(root, 0, 0)
	return ans
}

func dfsWidth(root *TreeNode, depth, pos int) {
	if root == nil {
		return
	}
	if _, ok := width[depth]; !ok {
		width[depth] = pos
	}
	//每一层的宽度
	ans = max(ans, pos-width[depth]+1)
	dfsWidth(root.Left, depth+1, 2*pos)
	dfsWidth(root.Right, depth+1, 2*pos+1)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 给定一个无重复元素的有序整数数组 nums 。

// 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

// 列表中的每个区间范围 [a,b] 应该按如下格式输出：

// "a->b" ，如果 a != b
// "a" ，如果 a == b

// 输入：nums = [0,1,2,4,5,7]
// 输出：["0->2","4->5","7"]
// 解释：区间范围是：
// [0,2] --> "0->2"
// [4,5] --> "4->5"
// [7,7] --> "7"

func summaryRanges(nums []int) []string {
	var sb []string
	var i int
	for j := 0; j < len(nums); j++ {
		if j+1 == len(nums) || nums[j]+1 != nums[j+1] {
			var s string
			s = s + strconv.Itoa(nums[i])
			if i != j {
				s = s + "->" + strconv.Itoa(nums[j])
			}
			sb = append(sb, s)
			i = j + 1
		}
	}
	return sb
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	left := &TreeNode{
		Val: 2,
	}
	root.Left = left

	fmt.Println(rightSideView(root))
}
