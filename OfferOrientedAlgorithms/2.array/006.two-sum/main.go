package main

import "fmt"

/*
题目：
输入一个递增排序的数组和一个值k，请问如何在数组中找出两个和为k的数字并返回它们的下标？
假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。
例如，输入数组[1，2，4，6，10]，k的值为8，数组中的数字2与6的和为8，它们的下标分别为1与3。

https://leetcode-cn.com/problems/kLl5u1/
*/

/*
笔记：
排序数组：二分查找，双指针
*/

func main() {
	fmt.Println(twoSum([]int{1, 2, 4, 6, 10}, 8))
}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		v := numbers[i] + numbers[j]
		if v == target {
			return []int{i, j}
		} else if v < target {
			i++
		} else {
			j--
		}
	}
	return nil
}
