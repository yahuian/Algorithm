package _0001

/*
题目：
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标。

https://leetcode-cn.com/problems/two-sum/
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	var res []int
	for i, v := range nums {
		k := target - v
		if val, ok := m[k]; !ok { // 用差值去查询 map
			m[v] = i // 存储的时候存原始值
		} else {
			res = []int{val, i}
		}
	}
	return res
}
