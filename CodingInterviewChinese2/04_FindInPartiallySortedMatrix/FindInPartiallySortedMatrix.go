package _04

// 面试题4：二维数组中的查找
// 题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按
// 照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个
// 整数，判断数组中是否含有该整数。

// 解题要点：充分利用题目提供的条件，
// 每一行都按照从左到右递增的顺序排序，
// 每一列都按照从上到下递增的顺序排序
func FindInPartiallySortedMatrix(nums [][]int, key int) bool {
	rows := len(nums)
	if rows == 0 {
		return false
	}
	column := len(nums[0])
	if column == 0 {
		return false
	}

	// 从矩阵右上角开始
	for i, j := 0, column-1; i < rows && j >= 0; {
		if nums[i][j] == key {
			return true
		} else if nums[i][j] > key {
			j--
			continue
		} else {
			i++
		}
	}
	return false
}
