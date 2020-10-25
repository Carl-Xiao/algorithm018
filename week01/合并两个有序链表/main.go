package main

// 示例：

// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4

//1 利用新的数组存放元素，然后再构建链表 时间复杂度为o(n) 空间复杂度o(n)
//2 使用递归  由于是有序的，这个时候只需要修改链表的next的指针值即可
//3 使用多余的存储空间
type ListNode struct {
	Val  int
	Next *ListNode
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	if l1.Val < l2.Val {
// 		l1.Next = mergeTwoLists(l1.Next, l2)
// 		return l1
// 	} else {
// 		l2.Next = mergeTwoLists(l1, l2.Next)
// 		return l2
// 	}
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tempNode := &ListNode{}
	prev := &ListNode{}
	prev = tempNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tempNode.Next = l1
			l1 = l1.Next
		} else {
			tempNode.Next = l2
			l2 = l2.Next
		}
		tempNode = tempNode.Next
	}
	//容易遗忘的点,上面的跳出循环的条件是某个链表结束 所以剩下的节点一定是大于当前的所有点，直接拼接在后面即可
	if l1 == nil {
		tempNode.Next = l2
	} else {
		tempNode.Next = l1
	}
	return prev.Next
}
func printNode(l1 *ListNode) {
	for l1 != nil {
		println(l1.Val)
		l1 = l1.Next
	}
}

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 4}
	l1.Next = l2
	l2.Next = l3

	l4 := &ListNode{Val: 1}
	l5 := &ListNode{Val: 3}
	l6 := &ListNode{Val: 4}

	l4.Next = l5
	l5.Next = l6

	printNode(mergeTwoLists(l1, l4))
}
