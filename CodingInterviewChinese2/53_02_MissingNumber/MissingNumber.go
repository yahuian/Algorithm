package _53_02

// 面试题53（二）：0到n-1中缺失的数字
// 题目：一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字
// 都在范围0到n-1之内。在范围0到n-1的n个数字中有且只有一个数字不在该数组
// 中，请找出这个数字。

// 转换为寻找第一个下标和值不相等的数字
func MissingNumber(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	i, j := 0, len(arr)-1
	for i <= j {
		mid := (i + j) / 2
		if arr[mid] == mid {
			i = mid + 1
		} else {
			if mid == 0 || arr[mid-1] == mid-1 {
				return mid
			} else {
				j = mid - 1
			}
		}
	}

	// 缺失的元素是最后一位
	if i == len(arr) {
		return len(arr)
	}

	return -1
}
