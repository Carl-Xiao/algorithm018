package main

//TreeNode 根节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据一棵树的前序遍历与中序遍历构造二叉树。

// 注意:
// 你可以假设树中没有重复的元素。

// 例如，给出

// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]

//分析

//1 前序第一个点  就是根节点
//2 通过根节点将 将中序遍历的数组分成两部分 left right

//  3
// / \
// 9  20
//   /   \
//  15    7
//第一次  从前序找到根节点 1  中序 left 4 2 5  right 3

//所以问题变成了 如何合理拆分
//left  前序 start,start+左子树的长度  中序 start ,根节点-1
//right 前序 start+左子树的长度+1 end  中序 根节点+1 end

func buildTree(preorder []int, inorder []int) *TreeNode {
	ma := make(map[int]int)

	for i := 0; i < len(inorder); i++ {
		ma[inorder[i]] = i
	}

	var a func(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, ma map[int]int) *TreeNode

	a = func(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, inMap map[int]int) *TreeNode {
		//跳出递归的条件 前序或者中序的整个数据长度遍历完
		if preStart > preEnd || inStart > inEnd {
			return nil
		}
		root := &TreeNode{Val: preorder[preStart]}
		inorderRoot := ma[root.Val]
		//左子树的长度
		leftTreeLen := inorderRoot - inStart

		root.Left = a(preorder, preStart+1, preStart+leftTreeLen, inorder, inStart, inorderRoot-1, inMap)
		//preStart+leftTreeLen+1 有点不懂先记录着回看再想明白
		root.Right = a(preorder, preStart+leftTreeLen+1, preEnd, inorder, inorderRoot+1, inEnd, inMap)

		return root
	}

	return a(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, ma)
}
