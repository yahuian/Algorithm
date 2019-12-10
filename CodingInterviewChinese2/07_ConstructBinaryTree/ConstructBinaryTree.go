package _07

import . "Algorithm/DataStructure"

// 面试题7：重建二叉树
// 题目：输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输
// 入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,
// 2, 4, 7, 3, 5, 6, 8}和中序遍历序列{4, 7, 2, 1, 5, 3, 8, 6}，则重建出
// 图2.6所示的二叉树并输出它的头结点。

// 递归法
func Construct(pre, in []int) *TreeNode {
	if len(pre) == 0 && len(in) == 0 {
		return nil
	}
	root := &TreeNode{Val: pre[0]}
	var leftTreePre, leftTreeIn []int   // 左子树的前序和中序
	var rightTreePre, rightTreeIn []int // 右子树的前序和中序

	for i := 0; i < len(in); i++ {
		if in[i] == root.Val {
			if i != 0 { // 左子树不为空
				leftTreePre = pre[1 : i+1]
				leftTreeIn = in[0:i]
			}
			if i != len(in)-1 { // 右子树不为空
				rightTreePre = pre[i+1:]
				rightTreeIn = in[i+1:]
			}
		}
	}

	root.Left = Construct(leftTreePre, leftTreeIn)
	root.Right = Construct(rightTreePre, rightTreeIn)
	return root
}

// 测试
func myTest(preTree, inTree []int) [][]int {
	root := Construct(preTree, inTree)
	pre := PreOrder(root)
	in := InOrder(root)
	post := PostOrder(root)
	return [][]int{pre, in, post}
}
