package _0191

func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		count++
		num = num & (num - 1) // 消除num最低位的1
	}
	return count
}
