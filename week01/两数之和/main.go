package main

//两数之和  暴力求解，两重for循环
//利用hash结构来做  nums的value为键，index为键

//参考官方推荐的例子
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {

}
