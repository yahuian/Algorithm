package _0257

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var result []string

	var preOrder func(*TreeNode, string)
	preOrder = func(root *TreeNode, path string) {
		if root != nil {
			value := root.Val
			path += strconv.Itoa(value)
			if root.Left == nil && root.Right == nil {
				result = append(result, path)
			} else {
				path += "->"
				preOrder(root.Left, path)
				preOrder(root.Right, path)
			}
		}
	}

	preOrder(root, "")
	return result
}
