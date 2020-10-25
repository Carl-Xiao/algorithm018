package main

// 示例 1:

// 输入: [1,2,3]
// 输出: [1,2,4]
// 解释: 输入数组表示数字 123。

// 如同10进制的加减法一样  反着处理应该更好
//
func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		//参考  如果为0 ,则直接返回即可
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}
