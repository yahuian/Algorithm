package _0152

import "math"

// 暴力求解法
//func maxProduct(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	max := nums[0]
//	for i := 0; i < len(nums); i++ {
//		tmp := 1
//		tmp *= nums[i]
//		if tmp > max {
//			max = tmp
//		}
//		for j := i + 1; j < len(nums); j++ {
//			tmp *= nums[j]
//			if tmp > max {
//				max = tmp
//			}
//		}
//	}
//	return max
//}

// 动态规划
// https://leetcode-cn.com/problems/maximum-product-subarray/solution/hua-jie-suan-fa-152-cheng-ji-zui-da-zi-xu-lie-by-g/
func maxProduct(nums []int) int {
	// MAX是最终的结果
	// iMax记录的是当前的最大值
	// iMin记录的是当前的最小值
	MAX, iMax, iMin := math.MinInt64, 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			iMax, iMin = iMin, iMax
		}
		iMax = max(nums[i], iMax*nums[i])
		iMin = min(nums[i], iMin*nums[i])

		MAX = max(iMax, MAX)
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
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
