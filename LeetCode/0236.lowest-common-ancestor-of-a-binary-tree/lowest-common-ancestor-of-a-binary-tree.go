package _0236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findPorQ(root, p, q)
}

func findPorQ(root, p, q *TreeNode) *TreeNode {
	switch root {
	case nil:
		return nil
	case p:
		return p
	case q:
		return q
	}
	left := findPorQ(root.Left, p, q)
	right := findPorQ(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}
