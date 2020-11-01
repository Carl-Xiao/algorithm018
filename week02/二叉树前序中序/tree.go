package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历  left-root-right
func inorderTraversal(root *TreeNode) (result []int) {
	var order func(root *TreeNode)

	order = func(root *TreeNode) {
		if root != nil {
			order(root.Left)
			result = append(result, root.Val)
			order(root.Right)

		}
	}
	order(root)
	return nil
}

//前序遍历  root-left-right
func preorderTraversal(root *TreeNode) (result []int) {
	var order func(root *TreeNode)

	order = func(root *TreeNode) {
		if root != nil {
			result = append(result, root.Val)
			order(root.Left)
			order(root.Right)
		}
	}
	order(root)
	return nil

}

// class Solution {

//     private List<List<Integer>> result = new ArrayList<>();

//     public List<List<Integer>> levelOrder(Node root) {
//         if (root != null) traverseNode(root, 0);
//         return result;
//     }

//     private void traverseNode(Node node, int level) {
//         if (result.size() <= level) {
//             result.add(new ArrayList<>());
//         }
//         result.get(level).add(node.val);
//         for (Node child : node.children) {
//             traverseNode(child, level + 1);
//         }
//     }
// }
