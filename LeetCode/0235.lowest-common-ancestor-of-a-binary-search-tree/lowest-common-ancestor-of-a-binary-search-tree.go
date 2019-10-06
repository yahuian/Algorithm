package _0235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findPorQ(root, p, q)
}

func findPorQ(root, p, q *TreeNode) *TreeNode {
	// p,q都在右子树
	if root.Val < p.Val && root.Val < q.Val {
		return findPorQ(root.Right, p, q)
	}
	// p,q都在左子树
	if root.Val > p.Val && root.Val > q.Val {
		return findPorQ(root.Left, p, q)
	}
	// p,q分散在左右子树中
	return root
}
