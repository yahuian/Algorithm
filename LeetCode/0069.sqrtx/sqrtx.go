package _0069

// 通过二分法逐步逼近left和right
// 当两者的整数部分相同时结束循环
func mySqrt(x int) int {
	var left float64
	right := float64(x)
	mid := (left + right) / 2
	for int(left) < int(right) {
		square := mid * mid
		if square > float64(x) {
			right = mid
		} else {
			left = mid
		}
		mid = (left + right) / 2
	}
	return int(mid)
}
