package main

import "fmt"

// 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

// 说明：

// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

//有序数组
// 1 暴力破解法 使用多余的数组空间存放，然后再进行排序
// 2 从小到大的数组  最优解肯定是从后往前处理数据
// 利用原数组空间修改数据 而不是移动
func merge(nums1 []int, m int, nums2 []int, n int) {
	end1 := m - 1
	end2 := n - 1
	l := m + n - 1
	for end1 >= 0 && end2 >= 0 {
		if nums1[end1] < nums2[end2] {
			nums1[l] = nums2[end2]
			end2--
		} else {
			nums1[l] = nums1[end1]
			end1--
		}
		l--
	}
	//碰到 [0] 1 [1] 1单元测试不通过
	for end2 >= 0 {
		nums1[l] = nums2[end2]
		l--
		end2--
	}
	fmt.Print(nums1)
}

func main() {
	//测试还不知道咋个传值 merge([]int{0}, 1, []int{1}, 1)
}
