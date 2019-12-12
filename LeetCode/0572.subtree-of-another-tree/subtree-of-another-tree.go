package _0572

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	result := false

	if s != nil && t != nil {
		if s.Val == t.Val {
			result = isSameTree(s, t)
		}
		if !result {
			result = isSubtree(s.Left, t)
		}
		if !result {
			result = isSubtree(s.Right, t)
		}
	}

	return result
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	result := false

	if s != nil && t != nil {
		if s.Val == t.Val {
			result = isSameTree(s.Left, t.Left)
		}

		if result {
			result = isSameTree(s.Right, t.Right)
		}
	}

	return result
}
