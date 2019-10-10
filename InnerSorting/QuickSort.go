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
	pivot := s[left]
	i, j := left, right
	for i < j {
		// 从右向左找小于pivot的数
		for s[j] >= pivot && i < j {
			j--
		}
		// 从左向右找大于pivot的数
		for s[i] <= pivot && i < j {
			i++
		}
		// 交换arr[i]和arr[j]的位置
		s[i], s[j] = s[j], s[i]
	}
	// pivot归位
	s[i], s[left] = s[left], s[i]
	return i
}

// 参考资料：https://bbs.codeaha.com/thread-4419-1-1.html
