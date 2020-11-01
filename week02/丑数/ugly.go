package main

// 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

//这个题有点难理解，不过只要理解题意做还是挺方便的 2* 3* 5* 第n个数一定是这些之中的一个，并且会排序从小到大
//定义一个数组存放所有的丑数
//所以这就是动态规划？
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func nthUglyNumber(n int) int {
	//创建3个索引指针
	var a, b, c int
	result := make([]int, n, n)
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = min(result[c]*5, min(result[a]*2, result[b]*3))

		//每次都选择最小的值
		if result[i] == result[a]*2 {
			a++
		}
		if result[i] == result[b]*3 {
			b++
		}
		if result[i] == result[c]*5 {
			c++
		}

	}
	return result[n-1]
}

func main() {

}
