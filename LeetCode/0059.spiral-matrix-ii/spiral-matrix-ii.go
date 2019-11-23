package _0059

func generateMatrix(n int) [][]int {
	up := 0
	down := n - 1
	left := 0
	right := n - 1
	temp := 1
	var answer [][]int
	for i := 0; i < n; i++ {
		inline := make([]int, n)
		answer = append(answer, inline)
	}
	for {
		// 向右移动，直至最右
		for i := left; i <= right; i++ {
			answer[up][i] = temp
			temp++
		}
		// 重新设定上边界，若上边界大于下边界，则遍历完成
		up++
		if up > down {
			break
		}

		// 向下移动，直至最底
		for i := up; i <= down; i++ {
			answer[i][right] = temp
			temp++
		}
		// 重新设定右边界，若右边界小于左边界，则遍历完成
		right--
		if right < left {
			break
		}

		// 向左移动，直至最左
		for i := right; i >= left; i-- {
			answer[down][i] = temp
			temp++
		}
		// 重新设定下边界，若下边界大于上边界，则遍历完成
		down--
		if down < up {
			break
		}

		// 向上移动，直至最上
		for i := down; i >= up; i-- {
			answer[i][left] = temp
			temp++
		}
		// 重新设定左边界，若左边界大于右边界，则遍历完成
		left++
		if left > right {
			break
		}
	}
	return answer
}
