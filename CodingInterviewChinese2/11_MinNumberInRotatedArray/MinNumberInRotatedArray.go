package _11

// 面试题11：旋转数组的最小数字
// 题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如数组
// {3, 4, 5, 1, 2}为{1, 2, 3, 4, 5}的一个旋转，该数组的最小值为1。

func MinNumberInRotatedArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	p1 := 0
	p2 := len(nums) - 1
	// 将mid设置为p1,一旦发现第一个数小于最后一个数，说明数组旋转了0位
	// 即没旋转，可以直接返回答案nums[p1],即第一个数
	mid := p1
	for nums[p1] >= nums[p2] {
		// 说明p1指向了第一个递增数组的尾部
		// p2指向了第二个递增数组的头部，即最小值
		if p2-p1 == 1 {
			mid = p2
			break
		}

		mid = (p1 + p2) / 2

		// 如果p1,mid,p2三者指向的数字相同，则没法判断最小值位于什么位置，只能遍历查找
		if nums[p1] == nums[mid] && nums[mid] == nums[p2] {
			res := nums[p1]
			for i := p1; i <= p2; i++ {
				if res > nums[i] {
					res = nums[i]
				}
			}
			return res
		}

		if nums[mid] >= nums[p1] { // 说明mid位于第一个递增数组
			p1 = mid
		} else if nums[mid] <= nums[p2] { // 说明mid位于第二个递增数组
			p2 = mid
		}

	}
	return nums[mid]
}
