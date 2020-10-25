package main

//https://leetcode-cn.com/problems/trapping-rain-water/

//这个题有点搞不懂get到题的核心
//所以第一次完成看答案去反查如何解决问题

// 先用自己好理解的双指针法,这点跟寻找最大的面积那道题不一样
// 我们得首先明白，木桶原理，当前容纳得面积得按照最小的处理

//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
func trap(height []int) int {
	var left, right = 0, len(height) - 1
	var result, leftMax, rightMax int
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				//设置左边最高柱子
				leftMax = height[left]
			} else {
				//右边必定有柱子挡水，所以，遇到所有值小于等于leftMax的，全部加入水池
				result += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				//设置右边最高柱子
				rightMax = height[right]
			} else {
				//左边必定有柱子挡水，所以，遇到所有值小于等于rightMax的，全部加入水池
				result += rightMax - height[right]
			}
			right--
		}
	}
	return result
}

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}
