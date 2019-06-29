/*
题目描述：
在一个二维数组中（每个一维数组的长度相同），
每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/
package main

import "fmt"

type s1 []int //一维数组
type s2 []s1  //由s1组成的二维数组

func main() {
	var s = s2{
		s1{1, 2, 8, 9},
		s1{2, 4, 9, 12},
		s1{4, 7, 10, 13},
		{6, 8, 11, 15}}

	fmt.Print(serach(s, 4, 4, 7))
}

func serach(s s2, row int, col int, key int) bool {
	//检测二维数组的合法性
	if s != nil && row > 0 && col > 0 {
		//每次都从数组的右上角开始判断
		var r = 0
		var c = col - 1
		for r < row && c >= 0 {
			if s[r][c] == key {
				//和key相等
				return true
			} else if s[r][c] > key {
				//大于key
				c-- //删除该列
			} else {
				//小于key
				r++ //删除改行
			}
		}
	}
	return false
}

/*
总结：
1.本题刚开始自然而然地想到的是暴力求解法，即将二维数组全部遍历一遍，
这样就完全没有利用题目中所给的条件。
2.本题体现了分类思想
*/
