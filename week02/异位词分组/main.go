package main

import "sort"

type bytes []byte

// 自定义排序规则
func (b bytes) Len() int {
	return len(b)
}
func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	mk := make(map[string]int)

	for _, str := range strs {
		//没有提供合适的方法，重写了
		temp := bytes(str)
		sort.Sort(temp) // 将字符串排序
		k := string(temp)
		if idx, ok := mk[k]; !ok {
			mk[k] = len(result)
			result = append(result, []string{str})
		} else {
			result[idx] = append(result[idx], str)
		}
	}
	return result
}

func main() {

}
