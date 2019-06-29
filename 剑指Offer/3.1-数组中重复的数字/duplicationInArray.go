package main

import (
	"fmt"
)

func main() {
	var s = []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(func3(s))
}

//方法一：暴力求解，两两直接对比
//或者先排序，再比对
func func1(s []int) int {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return s[i]
			}
		}
	}
	return -1 //表示该数组中无重复元素出现
}

//方法二：顺序遍历数组，每遍历一个就将其存放在map中
func func2(s []int) int {
	var tempMap = map[int]int{}
	for i := 0; i < len(s); i++ {
		for _, v := range tempMap {
			if v == s[i] {
				//map中有该值，即v是数组中重复的元素
				return v
			}
		}
		//map中无该值，将其添加进map
		tempMap[s[i]] = s[i]
	}
	return -1
}

//方法三,充分利用题目所给条件
func func3(s []int) int {
	var i int = 0
	for i < len(s) {
		if i == s[i] {
			//如果第i个位置的值正好和i相等，则检查数组下一项
			i++
		} else {
			//如果第i个位置的值m和i不相等，则检查s[i]和s[m]的情况
			var m int = s[i]
			if s[i] == s[m] {
				//数组中有重复数字
				return s[i]
			} else {
				//交换两者的位置
				s[i], s[m] = s[m], s[i]
			}
		}
	}
	return -1
}
