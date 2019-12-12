package _0101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return compare(root, root)
}

func compare(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return compare(l.Left, r.Right) && compare(l.Right, r.Left)
}
