package main

import (
	"fmt"
)

func main() {
	var myArray = []int{2, 3, 5, 4, 3, 2, 6, 7}
	fmt.Println(func1(myArray))
}

//方法一：新创建一个大小为len(s)的数组
func func1(s []int) int {
	var temp = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if temp[s[i]] == s[i] {
			return s[i]
		} else {
			temp[s[i]] = s[i]
		}
	}
	fmt.Println("输入的数组不符合题意")
	return -1
}
