package _0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-result)) {
				result = sum
			}
			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				return sum
			}
		}
	}
	return result
}

// 相比三数之和，少了一部分去重的代码
