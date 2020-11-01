package main

type Node struct {
	Val      int
	Children []*Node
}

//N叉树的前序遍历
//root->left-> right
//root->孩子
func preorder(root *Node) (result []int) {
	var order func(root *Node)

	order = func(root *Node) {
		if root != nil {
			result = append(result, root.Val)
			for _, value := range root.Children {
				order(value)
			}
		}
	}
	order(root)
	return nil

}
