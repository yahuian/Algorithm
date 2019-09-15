package _0102

import "Algorithm/DataStructure"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	aQueue := DataStructure.Queue{}
	result := [][]int{}
	// 初始化队列
	level := 0
	aQueue.Push(root)

	for aQueue.Len() != 0 {
		length := aQueue.Len()
		// 每增加一层，二维切片就添加一行
		result = append(result, []int{})
		for i := 0; i < length; i++ {
			node, _ := aQueue.Pop()
			result[level] = append(result[level], node.(*TreeNode).Val)

			// 将当前节点的左右孩子入队
			if node.(*TreeNode).Left != nil {
				aQueue.Push(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				aQueue.Push(node.(*TreeNode).Right)
			}
		}
		level++
	}
	return result
}
