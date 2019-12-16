package _32_02

import . "Algorithm/DataStructure"

// 面试题32（二）：分行从上到下打印二叉树
// 题目：从上到下按层打印二叉树，同一层的结点按从左到右的顺序打印，每一层
// 打印到一行。

func PrintTreesInLines(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	var result [][]int
	queue = append(queue, root)
	for len(queue) != 0 {
		var temp []int
		length := len(queue)
		// 控制按层输出
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			temp = append(temp, node.Val)
		}
		result = append(result, temp)
	}
	return result
}
