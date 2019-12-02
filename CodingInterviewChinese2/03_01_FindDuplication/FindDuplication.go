package _03_01

// 面试题3（一）：找出数组中重复的数字
// 题目：在一个长度为n的数组里的所有数字都在0到n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
// 也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数组{2, 3, 1, 0, 2, 5, 3}，
// 那么对应的输出是重复的数字2或者3。

// 方法1：排序数组
// 方法2：哈希表
// 方法3：充分利用题目条件-所有数字都在0到n-1的范围内
func FindDuplication(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	var res []int
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
			continue
		}
		val := nums[i]
		if nums[i] == nums[val] {
			res = append(res, nums[i])
		} else {
			nums[i], nums[val] = nums[val], nums[i]
		}
	}
	return res
}
