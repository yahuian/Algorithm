package _0230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	array := inOrder(root)
	if len(array) < k {
		return -1
	}
	return array[k-1]
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := inOrder(root.Left)
	val := root.Val
	right := inOrder(root.Right)
	return append(append(left, val), right...)
}
