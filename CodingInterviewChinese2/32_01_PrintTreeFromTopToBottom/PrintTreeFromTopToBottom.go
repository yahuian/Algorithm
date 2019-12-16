package _32_01

import (
	. "Algorithm/DataStructure"
)

// 面试题32（一）：不分行从上往下打印二叉树
// 题目：从上往下打印出二叉树的每个结点，同一层的结点按照从左到右的顺序打印。

func PrintTreeFromTopToBottom(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	var result []int
	// 头结点先入队列
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		result = append(result, node.Val)
		queue = queue[1:]
	}
	return result
}
