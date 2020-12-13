package main

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词

// 利用hashmap计数即可
func isAnagram(s string, t string) bool {
	mac := make(map[rune]int)
	for _, value := range s {
		mac[value]++
	}
	for _, value := range t {
		if temp, ok := mac[value]; ok {
			temp = temp - 1
			if temp == 0 {
				delete(mac, value)
			} else {
				mac[value] = temp
			}
		} else {
			return false
		}
	}
	return len(mac) == 0
}
