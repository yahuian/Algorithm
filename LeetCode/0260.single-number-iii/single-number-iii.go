package _0260

func singleNumber(nums []int) []int {
	if len(nums) < 2 {
		return nil
	}

	temp := 0
	for _, v := range nums {
		temp ^= v
	}

	bit1 := findFirstBit1(temp)

	var ans1, ans2 int

	for _, v := range nums {
		if isBit1(v, bit1) {
			ans1 ^= v
		} else {
			ans2 ^= v
		}
	}

	return []int{ans1, ans2}
}

// 找整数x的二进制表示中，最右边的1位置（从0计数）
// eg: 输入2(即0010)
//     输出 1
func findFirstBit1(x int) int {
	i := 0
	for x&1 == 0 {
		i++
		x = x >> 1
	}
	return i
}

// 判断啊x的二进制表示中，最右边第i位是不是1（从0计数）
func isBit1(x, i int) bool {
	x = x >> i
	return x&1 == 1
}

// <<剑指offer>> 56(1)
// https://github.com/zhedahht/CodingInterviewChinese2/blob/master/56_01_NumbersAppearOnce/NumbersAppearOnce.cpp
