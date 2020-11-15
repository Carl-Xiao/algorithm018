package main

import (
	"fmt"
)

// 127. 单词接龙

// 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

// 每次转换只能改变一个字母。
// 转换过程中的中间单词必须是字典中的单词。
// 说明:

// 如果不存在这样的转换序列，返回 0。
// 所有单词具有相同的长度。
// 所有单词只由小写字母组成。
// 字典中不存在重复的单词。
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

// 示例 1:

// 输入:
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]
//"hot", "dot", "dog", "lot", "log"

// 输出: 5

// 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//      返回它的长度 5。

// 耗时3小时+

//问题 1 单向递归被判超时,不知道是什么机制
//问题 2 对golang的使用出现了致命的失误

//问题3 当遇到这种测试的时候出现问题
//"ymain"
// "oecij"
// ["ymann","yycrj","oecij","ymcnj","yzcrj","yycij","xecij","yecij","ymanj","yzcnj","ymain"]

//总之在个题耗费 2020-11-15一下午的时间
func main() {
	fmt.Println(ladderLength("ymain", "oecij", []string{"ymann", "yycrj", "oecij", "ymcnj", "yzcrj", "yycij", "xecij", "yecij", "ymanj", "yzcnj", "ymain"}))

	// fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordmap := make(map[string]bool)
	for _, value := range wordList {
		wordmap[value] = true
	}
	if _, ok := wordmap[endWord]; !ok {
		return 0
	}

	visited := make(map[string]bool)
	beginqueue := new([]string)
	endqueue := new([]string)
	*beginqueue = append(*beginqueue, beginWord)
	*endqueue = append(*endqueue, endWord)

	if _, ok := wordmap[beginWord]; ok {
		delete(wordmap, beginWord)
	}

	count := 1
	for len(*beginqueue) > 0 && len(*endqueue) > 0 {
		res := find(count, beginqueue, endqueue, visited, wordmap)

		if res != -1 {
			return res
		}
		count++
	}
	return 0
}
func find(min int, beginqueue, endqueue *[]string, visited, dict map[string]bool) int {
	var count int
	pre := false
	start := len(*beginqueue)
	end := len(*endqueue)

	if start <= end {
		count = start
		pre = true
	} else {
		count = end
	}
	fmt.Println(*beginqueue, *endqueue, visited, count, pre)

	for i := 0; i < count; i++ {
		var value string
		if pre {
			value = (*beginqueue)[0]
			*beginqueue = (*beginqueue)[1:]
		} else {
			value = (*endqueue)[0]
			*endqueue = (*endqueue)[1:]
		}
		chs := []byte(value)
		for j := 0; j < len(chs); j++ {
			for c := 'a'; c <= 'z'; c++ {
				if byte(c) == chs[j] {
					continue
				}
				old := chs[j]
				chs[j] = byte(c)
				target := string(chs)
				if _, ok := dict[target]; ok {
					if pre {
						if contains(*endqueue, target) {
							// fmt.Println(pre, *endqueue, target)
							return min + 1
						}
						ok := visited[target]
						if !ok {
							*beginqueue = append(*beginqueue, target)
							visited[target] = true
						}

					} else {
						if contains(*beginqueue, target) {
							fmt.Println(*beginqueue, target)
							return min + 1
						}
						ok := visited[target]
						if !ok {
							*endqueue = append(*endqueue, target)
							visited[target] = true
						}

					}
				}

				chs[j] = old
			}
		}
	}
	return -1
}

func contains(word []string, taget string) bool {
	for _, value := range word {
		if value == taget {
			return true
		}
	}
	return false
}
