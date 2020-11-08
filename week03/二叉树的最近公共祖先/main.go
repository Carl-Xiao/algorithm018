// 236题目 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。
package main

//TreeNode 根节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//分析
// 从左+从右 一起找
// 假设节点分布在左边与右边,就返回当前root

// 如果排列在一边 则只返回q或者q的一个节点即可
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//边界情况
	if root == nil || root == p || root == q {
		return root
	}

	//进入下一层
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//处理逻辑
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}
