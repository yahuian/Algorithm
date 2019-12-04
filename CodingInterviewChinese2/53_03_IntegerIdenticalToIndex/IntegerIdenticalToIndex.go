package _53_03

// 面试题53（三）：数组中数值和下标相等的元素
// 题目：假设一个单调递增的数组里的每个元素都是整数并且是唯一的。请编程实
// 现一个函数找出数组中任意一个数值等于其下标的元素。例如，在数组{-3, -1,
// 1, 3, 5}中，数字3和它的下标相等。

func IntegerIdenticalToIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	i, j := 0, len(arr)-1
	for i <= j {
		mid := (i + j) / 2
		if mid == arr[mid] {
			return mid
		} else if arr[mid] > mid { // 如果值大于下标，则右边的所有值都大于下标，所以向左边查找
			j = mid - 1
		} else { // 如果值小于下标，则左边的所有值都小于下标，所以向右边查找
			i = mid + 1
		}
	}

	return -1
}
