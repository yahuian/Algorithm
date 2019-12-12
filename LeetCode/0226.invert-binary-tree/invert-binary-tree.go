package _0226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left

	if root.Left != nil {
		helper(root.Left)
	}

	if root.Right != nil {
		helper(root.Right)
	}
}
