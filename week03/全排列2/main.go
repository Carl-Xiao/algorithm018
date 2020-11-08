package main

//比上面多了一个判断是否重复的就可以了
func permuteUnique(nums []int) [][]int {
	result := &[][]int{}
	dfs(0, 1, nums, result)
	return *result
}

func dfs(start, k int, nums []int, results *[][]int) {
	if k == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*results = append(*results, tmp)
		return
	}
	used := make(map[int]int)
	for i := start; i < len(nums); i++ {
		if _, ok := used[nums[i]]; ok {
			continue
		}
		nums[i], nums[start] = nums[start], nums[i]
		dfs(start+1, k+1, nums, results)
		nums[i], nums[start] = nums[start], nums[i]
		used[nums[i]] = i
	}
}

func main() {

}
