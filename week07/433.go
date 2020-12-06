package main

import "fmt"

//跟单词接龙是一个意思 用bfs
func minMutation(start string, end string, bank []string) int {
	dic := make(map[string]bool)
	for _, v := range bank {
		dic[v] = true
	}
	if ok := dic[end]; !ok {
		return -1
	}

	ds := []string{"A", "C", "G", "T"}
	queue := []string{start}
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			word := queue[0]
			queue = queue[1:]
			for j := 0; j < len(word); j++ {
				for k := 0; k < len(ds); k++ {
					tmp := word[:j] + ds[k] + word[j+1:]
					if _, ok := dic[tmp]; ok {
						queue = append(queue, tmp)
						delete(dic, tmp)
					}
					if tmp == end {
						return step
					}
				}
			}
			size--
			step++
		}

	}
	return -1
}
func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{}))
}
