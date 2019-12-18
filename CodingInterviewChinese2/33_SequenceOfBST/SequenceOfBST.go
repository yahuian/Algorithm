package _33

// 面试题33：二叉搜索树的后序遍历序列
// 题目：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
// 如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。

func SequenceOfBST(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}
	root := nums[n-1]

	// 找左子树
	i := 0
	for ; i < n-1; i++ {
		if nums[i] > root {
			break
		}
	}

	// 判断右子树是否合规
	j := i
	for ; j < n-1; j++ {
		if nums[j] < root {
			return false
		}
	}

	// 递归
	left := SequenceOfBST(nums[0:i])
	right := SequenceOfBST(nums[i : n-1])

	return left && right
}
