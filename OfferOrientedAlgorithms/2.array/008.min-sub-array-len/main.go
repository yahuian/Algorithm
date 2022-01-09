package main

import "fmt"

func main() {
	fmt.Println(minSubArrayLen(7, []int{5, 1, 4, 3}))              // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
	fmt.Println(minSubArrayLen(1, []int{1}))                       // 1
	fmt.Println(minSubArrayLen(6, []int{10, 2, 3}))                // 1
}

/*
题目：
输入一个正整数组成的数组和一个正整数k，请问数组中和大于或等于k的连续子数组的最短长度是多少？
如果不存在所有数字之和大于或等于k的子数组，则返回0。
例如，输入数组[5，1，4，3]，k的值为7，和大于或等于7的最短连续子数组是[4，3]，因此输出它的长度2。
https://leetcode-cn.com/problems/2VG8Kg/
*/

/*
笔记：
使用双指针解决子数组之和的面试题有一个前提条件——数组中的所有数字都是正数。
如果数组中的数字有正数、负数和零，那么双指针的思路并不适用，
这是因为当数组中有负数时在子数组中添加数字不一定能增加子数组之和，
从子数组中删除数字也不一定能减少子数组之和。
*/

func minSubArrayLen(target int, nums []int) int {
	var res, i, j int
	for i < len(nums) && j < len(nums) && i <= j {
		sum := 0
		for k := i; k <= j; k++ {
			sum += nums[k]
		}
		if sum < target {
			j++
		} else {
			tmp := (j - i) + 1
			if res == 0 || tmp < res {
				res = tmp
			}
			i++
		}
	}
	return res
}
