package _0102

import "Algorithm/DataStructure"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先周游 BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	aQueue := DataStructure.Queue{}
	aQueue.Push(root)
	var res [][]int
	for aQueue.Len() != 0 {
		var tmp []int
		length := aQueue.Len()
		for i := 0; i < length; i++ {
			node, _ := aQueue.Pop()
			tmp = append(tmp, node.(*TreeNode).Val)
			if node.(*TreeNode).Left != nil {
				aQueue.Push(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				aQueue.Push(node.(*TreeNode).Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}

// 深度优先周游 DFS
func deepOrder(root *TreeNode) [][]int {
	var res [][]int
	helper(root, 0, &res)
	return res
}

// 注意res必须为指针类型，保证每次递归时的操作都时针对一个res
// 如果传递为[][]int类型，则每次递归时二维切片都会被复制
func helper(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	helper(root.Left, level+1, res)
	helper(root.Right, level+1, res)
}
