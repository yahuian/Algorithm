package _0007

import (
	"math"
)

func reverse(x int) int {
	var res int
	for x != 0 {
		pop := x % 10
		x /= 10
		res = res*10 + pop
		if math.Abs(float64(res)) > math.MaxInt32 {
			return 0
		}
	}
	return res
}
