package _0338

func countBits(num int) []int {
	list := make([]int, num+1)
	list[0] = 0
	for i := 1; i <= num; i++ {
		list[i] = list[i&(i-1)] + 1
	}
	return list
}
