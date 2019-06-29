package main

import (
	"fmt"
)

func main() {
	var s = []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(func1(s))
}

//方法一：暴力求解，两两直接对比
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
