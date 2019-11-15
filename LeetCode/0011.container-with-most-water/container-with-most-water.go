package _0011

// 双指针法
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	i, j := 0, len(height)-1
	maxCapacity := 0
	for i != j {
		capacity := min(height[i], height[j]) * (j - i)
		if capacity > maxCapacity {
			maxCapacity = capacity
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return maxCapacity
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
