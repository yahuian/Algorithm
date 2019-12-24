package _0111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	// 叶子节点是指没有子结点的节点
	if left == 0 {
		return right + 1
	}
	if right == 0 {
		return left + 1
	}

	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}
