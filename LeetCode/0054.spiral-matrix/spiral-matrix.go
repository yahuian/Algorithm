package _0054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var answer []int
	// 设定上下左右边界
	up := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for {
		// 向右移动，直至最右
		for i := left; i <= right; i++ {
			answer = append(answer, matrix[up][i])
		}
		// 重新设定上边界，若上边界大于下边界，则遍历完成
		up++
		if up > down {
			break
		}

		// 向下移动，直至最底
		for i := up; i <= down; i++ {
			answer = append(answer, matrix[i][right])
		}
		// 重新设定右边界，若右边界小于左边界，则遍历完成
		right--
		if right < left {
			break
		}

		// 向左移动，直至最左
		for i := right; i >= left; i-- {
			answer = append(answer, matrix[down][i])
		}
		// 重新设定下边界，若下边界大于上边界，则遍历完成
		down--
		if down < up {
			break
		}

		// 向上移动，直至最上
		for i := down; i >= up; i-- {
			answer = append(answer, matrix[i][left])
		}
		// 重新设定左边界，若左边界大于右边界，则遍历完成
		left++
		if left > right {
			break
		}

	}
	return answer
}

// https://leetcode-cn.com/problems/spiral-matrix/solution/cxiang-xi-ti-jie-by-youlookdeliciousc-3/
