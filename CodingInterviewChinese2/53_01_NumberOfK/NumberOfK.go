package _53_01

// 面试题53（一）：数字在排序数组中出现的次数
// 题目：统计一个数字在排序数组中出现的次数。例如输入排序数组{1, 2, 3, 3,
// 3, 3, 4, 5}和数字3，由于3在这个数组中出现了4次，因此输出4。

// 二分法的灵活应用
// O(logn)
func GetNumberOfK(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	i := getFirstK(arr, k)
	j := getLastK(arr, k)

	if i != -1 && j != -1 {
		return j - i + 1
	}
	return 0
}

// 找K第一次出现的位置
func getFirstK(arr []int, k int) int {
	i, j := 0, len(arr)-1
	for i <= j {
		mid := (i + j) / 2
		if arr[mid] == k {
			// 比正常的二分多了一步判断
			if mid-1 < 0 || arr[mid-1] != k {
				return mid
			} else {
				j = mid - 1
			}
		} else if arr[mid] > k {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}

// 找k最后一次出现的位置
func getLastK(arr []int, k int) int {
	i, j := 0, len(arr)-1
	for i <= j {
		mid := (i + j) / 2
		if arr[mid] == k {
			if mid+1 > len(arr)-1 || arr[mid+1] != k {
				return mid
			} else {
				i = mid + 1
			}
		} else if arr[mid] > k {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}
