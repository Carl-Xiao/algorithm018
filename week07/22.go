package main

func generateParenthesis(n int) []string {
	result := new([]string)

	dfs("", 2*n, n, n, result)

	return *result
}
func dfs(buf string, length, left, right int, result *[]string) {
	//移除条件,当长度等于2*n的时候，代表当前已经添加完毕
	if len(buf) == length {
		*result = append(*result, buf)
		return
	}
	//任意时候都可以
	if left > 0 {
		dfs(buf+"(", length, left-1, right, result)
	}

	if left < right {
		dfs(buf+")", length, left, right-1, result)
	}
}
