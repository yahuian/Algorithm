package _0008

import (
	"math"
)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}

	numMap := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	result := 0
	flag := 1
	i := 0

	// 首先消除str前面的空格
	for i < len(str) && string(str[i]) == " " {
		i++
	}

	// 判断正负
	// 这儿的代码有点奇怪主要是因为-+1这种情况
	if i < len(str) && string(str[i]) == "-" {
		flag = -1
	}
	if i < len(str) && (string(str[i]) == "+" || string(str[i]) == "-") {
		i++
	}

	// 处理剩余部分
	for ; i < len(str); i++ {
		// 判断是否是0-9的数字
		if value, ok := numMap[string(str[i])]; ok {
			result = result*10 + value
			// 判断是否溢出
			if result > math.MaxInt32 {
				if flag > 0 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}

		} else {
			break // 如果出现非数字的字符，则截止
		}
	}

	// 没有溢出的情况
	if flag > 0 {
		return result
	} else {
		return -1 * result
	}
}
