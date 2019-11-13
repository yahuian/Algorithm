package _0237

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 这道题比《剑指offer》中的18_01_DeleteNodeInList要简单一些
// 因为题干中说了，给定的是非尾巴的节点，所以要少一些判断条件
