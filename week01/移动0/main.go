package main

//必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。

//以前可能会想用额外的数组处理

// 示例:

// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		//犯了个错，没考虑到负数情况
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}

func main() {

}
