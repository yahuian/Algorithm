package _0015

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		// 为了去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if nums[l]+nums[r] > target {
				r--
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				temp := []int{nums[i], nums[l], nums[r]}
				res = append(res, temp)
				// 也是为了去重
				l, r = next(nums, l, r)
			}
		}
	}
	return res
}
func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}
	return l, r
}
