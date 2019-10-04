package _0098

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var list []int
	inOrder(root, &list) // 中序遍历后的二叉搜索树是有序的
	if len(list) == 1 {
		return true
	}
	for i, j := 0, 1; j < len(list); i, j = i+1, j+1 {
		if list[i] >= list[j] {
			return false
		}
	}
	return true
}

// 中序遍历二叉树
// leetcode中使用全局变量测试可能会有错，所以此处把list传入中序遍历函数
func inOrder(root *TreeNode, list *[]int) {
	if root != nil {
		inOrder(root.Left, list)
		*list = append(*list, root.Val)
		inOrder(root.Right, list)
	}
}
