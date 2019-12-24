package _0110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, res := helper(root)
	return res
}

// 后序遍历 左右中
func helper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, left := helper(root.Left)
	rightDepth, right := helper(root.Right)

	diff := leftDepth - rightDepth
	if left && right && diff >= -1 && diff <= 1 {
		return max(leftDepth, rightDepth) + 1, true
	}
	return 0, false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
