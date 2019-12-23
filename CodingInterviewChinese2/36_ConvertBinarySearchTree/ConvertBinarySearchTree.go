package _36

import . "Algorithm/DataStructure"

// 面试题36：二叉搜索树与双向链表
// 题目：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求
// 不能创建任何新的结点，只能调整树中结点指针的指向。

func Convert(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	// 将左子树构造成双向链表，并返回链表头节点
	left := Convert(root.Left)
	// 定位到左子树双向链表最后一个节点
	p := left
	for p != nil && p.Right != nil {
		p = p.Right
	}
	// 如果左子树双向链表不为空，将root追加到left后面
	if left != nil {
		p.Right = root
		root.Left = p
	}
	// 将右子树构造成双向链表，并返回链表头节点
	right := Convert(root)
	// 如果右子树双向链表不为空，将right追加到root后面
	if right != nil {
		right.Left = root
		root.Right = right
	}
	// 返回双向链表头
	if left != nil {
		return left
	} else {
		return root
	}
}

// https://www.nowcoder.com/questionTerminal/947f6eb80d944a84850b0538bf0ec3a5
