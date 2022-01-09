package main

import "fmt"

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}

/*
题目：
输入一个正整数组成的数组和一个正整数k，请问数组中和大于或等于k的连续子数组的最短长度是多少？
如果不存在所有数字之和大于或等于k的子数组，则返回0。
例如，输入数组[5，1，4，3]，k的值为7，和大于或等于7的最短连续子数组是[4，3]，因此输出它的长度2。
https://leetcode-cn.com/problems/ZVAVXX/
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	var res, i, j int
	for i < len(nums) && j < len(nums) && i <= j {
		p := 1
		for k := i; k <= j; k++ {
			p *= nums[k]
		}
		if p < k {
			// 10 子数组为 [10]
			// 10,5 子数组为 [5] [10,5]
			// 10,5,2 子数组为 [2] [5,2] [10,5,2]
			// 找到其规律为 j-i+1
			res += j - i + 1
			j++
		} else {
			i++
		}
	}
	return res
}
