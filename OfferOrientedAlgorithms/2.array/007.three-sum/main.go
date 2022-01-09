package main

import (
	"fmt"
	"sort"
)

/*
题目：
输入一个数组，如何找出数组中所有和为0的3个数字的三元组？
需要注意的是，返回值中不得包含重复的三元组。
例如，在数组[-1，0，1，2，-1，-4]中有两个三元组的和为0，它们分别是[-1，0，1]和[-1，-1，2]。
https://leetcode-cn.com/problems/1fGaJU/
*/

/*
笔记：
通过对原始数组排序，将三数之和转化为固定一个数字 i，然后在排序数组中查找和为 -i 的两数之和问题，
然后再考虑去重的逻辑，由简到难，逐步解决问题。TODO
排序：O(nlog n)；固定数组中每一个数字：O(n^2)；排序数组两数之和：O(n)；总时间复杂度为：O(n^2)
*/

func main() {
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	// 排序
	sort.Ints(nums)

	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		// 三数变两数
		r1, r2, ok := twoSum(nums[i+1:], nums[i])
		if !ok {
			continue
		}

		res = append(res, []int{nums[i], r1, r2})
	}

	return res
}

func twoSum(nums []int, target int) (int, int, bool) {
	i, j := 0, len(nums)-1
	for i < j {
		v := nums[i] + nums[j] + target
		if v == 0 {
			return nums[i], nums[j], true
		} else if v < 0 {
			i++
		} else {
			j--
		}

	}
	return 0, 0, false
}
