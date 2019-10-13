package _0111

import "Algorithm/DataStructure"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	aQueue := DataStructure.Queue{}
	aQueue.Push(root)
	var res [][]int
	for aQueue.Len() != 0 {
		length := aQueue.Len()
		var tmp []int
		for i := 0; i < length; i++ {
			node, _ := aQueue.Pop()
			tmp = append(tmp, node.(*TreeNode).Val)
			if node.(*TreeNode).Left != nil {
				aQueue.Push(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				aQueue.Push(node.(*TreeNode).Right)
			}
			// 左右孩子都为空
			if node.(*TreeNode).Left == nil && node.(*TreeNode).Right == nil {
				return len(res) + 1
			}
		}
		res = append(res, tmp)
	}
	return len(res)
}
