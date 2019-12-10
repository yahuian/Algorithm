package DataStructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

// 将[]int按照广度优先转化为一颗二叉树
// [1,2,3,4,NULL,5] -->      1
//				  		 2       3
//              	   4  NULL 5
func SliceToBinaryTree(s []int) *TreeNode {
	n := len(s)
	if n == 0 {
		return nil
	}

	queue := make([]*TreeNode, 1, n)
	root := &TreeNode{Val: s[0]}
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && s[i] != NULL {
			left := &TreeNode{Val: s[i]}
			node.Left = left
			queue = append(queue, left)
		}
		i++

		if i < n && s[i] != NULL {
			right := &TreeNode{Val: s[i]}
			node.Right = right
			queue = append(queue, right)
		}
		i++
	}
	return root
}

// -----------------递归版--------------------
// 前序遍历二叉树（根左右）
func PreOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	tmp := []int{root.Val}
	left := PreOrder(root.Left)
	right := PreOrder(root.Right)

	return append(append(tmp, left...), right...)
}

// 中序遍历（左根右）
func InOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := InOrder(root.Left)
	tmp := []int{root.Val}
	right := InOrder(root.Right)

	return append(append(left, tmp...), right...)
}

// 后序遍历（左右根）
func PostOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := PostOrder(root.Left)
	right := PostOrder(root.Right)
	tmp := []int{root.Val}

	return append(append(left, right...), tmp...)
}

// 测试
func myTest(s []int) [][]int {
	root := SliceToBinaryTree(s)
	pre := PreOrder(root)
	in := InOrder(root)
	post := PostOrder(root)
	return [][]int{pre, in, post}
}
