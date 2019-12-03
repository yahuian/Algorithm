package InnerSorting

// 归并排序
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	mid := length / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

// 合并两个有序数组为一个有序数组
func merge(l, r []int) []int {
	var res []int
	for len(l) != 0 && len(r) != 0 {
		if l[0] < r[0] {
			res = append(res, l[0])
			l = l[1:]
		} else {
			res = append(res, r[0])
			r = r[1:]
		}
	}

	if len(l) != 0 {
		res = append(res, l...)
	}

	if len(r) != 0 {
		res = append(res, r...)
	}

	return res
}
