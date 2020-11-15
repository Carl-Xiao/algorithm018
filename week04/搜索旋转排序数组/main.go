package main

import "fmt"

//33. 搜索旋转排序数组

// 给你一个整数数组 nums ，和一个整数 target 。

// 该整数数组原本是按升序排列，但输入时在预先未知的某个点上进行了旋转。（例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] ）。

// 请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

//二分查找
//输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4

//代码有些冗余,但是自己能很好理解
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {

		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		fmt.Println(left, right)

		//按照当前例子,前半段是正常的 后半段是异常
		if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1
}

func main() {
	fmt.Println(search([]int{5, 1, 3}, 5))
}
