package _0136

func singleNumber(nums []int) int {
	temp := 0
	for _, v := range nums {
		temp ^= v
	}
	return temp
}

// 位运算
// a ^ 0 = a
// a ^ a = 0
