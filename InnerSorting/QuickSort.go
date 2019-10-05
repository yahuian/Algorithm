package InnerSorting

func QuickSort(s []int) []int {
	return quickSort(s, 0, len(s)-1)
}

func quickSort(s []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(s, left, right)
		quickSort(s, left, partitionIndex-1)
		quickSort(s, partitionIndex+1, right)
	}
	return s
}

func partition(s []int, left, right int) int {
	i := left + 1
	j := right
	pivot := s[left] // 基准
	for i != j {
		// 从右边找小于基准的值
		for s[j] >= pivot && i < j {
			j--
		}
		// 从左边找大于基准的值
		for s[i] <= pivot && i < j {
			i++
		}
		// 大于基准的和小于基准的交换位置
		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}
	// i = j时会跳出外层for循环
	if s[i] < s[left] {
		s[i], s[left] = s[left], s[i]
		return i
	}
	return left
}

// 参考资料：https://bbs.codeaha.com/thread-4419-1-1.html
