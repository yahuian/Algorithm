package _0121

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt64
	max := 0
	for i := 0; i < len(prices); i++ {
		// 先找最小峰
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max { // 接着找最大峰
			max = prices[i] - min
		}
	}
	return max
}
