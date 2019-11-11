package _0053

import "math"

func maxSubArray(nums []int) int {
	// iMax表示以当前节点为终结节点的最大连续子序列和
	MAX, iMax := math.MinInt64, 0
	for i := 0; i < len(nums); i++ {
		iMax = max(nums[i]+iMax, nums[i])
		MAX = max(MAX, iMax)
	}
	return MAX
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
