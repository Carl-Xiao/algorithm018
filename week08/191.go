package main

//用位运算
func hammingWeight(num uint32) int {
	sum := 0
	for num > 0 {
		num &= num - 1
		sum++
	}
	return sum
}
