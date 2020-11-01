package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return nil
	}
	dfs(root, 0, &res)
	return res
}

//题目看着挺吓人,不过记住规则就好 深度遍历
func dfs(root *Node, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	for _, n := range root.Children {
		dfs(n, level+1, res)
	}
}

func main() {

}
