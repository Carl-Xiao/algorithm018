package main

//什么是异位词
//输入是通过将 s 的字母重新排列成 t 来生成变位词。
// 因此，如果 t 是 s 的变位词，对两个字符串进行排序将产生两个相同的字符串。
//Map解决 又快又好用
func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, value := range s {
		m[value]++
	}
	for _, value := range t {
		m[value]--
	}
	return len(m) == 0
}

func main() {

}
