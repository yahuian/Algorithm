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
