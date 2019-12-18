package _34

import . "Algorithm/DataStructure"

// 面试题34：二叉树中和为某一值的路径
// 题目：输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所
// 有路径。从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。

func PathInTree(root *TreeNode, key int) [][]int {
	var result [][]int

	// 函数类型的变量
	var preOrder func(root *TreeNode, path []int)
	preOrder = func(root *TreeNode, path []int) {
		if root != nil {
			path = append(path, root.Val)
			// 到叶子结点了，为一条路径
			if root.Left == nil && root.Right == nil {
				result = append(result, path)
			} else {
				preOrder(root.Left, path)
				preOrder(root.Right, path)
			}
		}
	}

	preOrder(root, []int{})
	// 此时result中为二叉树所有的路径

	var res [][]int

	// 计算路径和是否和key相等
	for _, path := range result {
		var temp int
		for _, v := range path {
			temp += v
		}
		if temp == key {
			res = append(res, path)
		}
	}

	return res
}
