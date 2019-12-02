package _0905

// 双指针法
func sortArrayByParity(A []int) []int {
	i, j := 0, len(A)-1
	for i < j {
		// 向后移动，直到它指向奇数
		for A[i]%2 == 0 && i < j {
			i++
		}
		// 向前移动，直到它指向偶数
		for A[j]%2 != 0 && i < j {
			j--
		}
		// 奇偶交换
		if i < j {
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
