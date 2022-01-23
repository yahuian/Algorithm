package _0011

/*
面积 = min(h[i],h[j]) * (j-i)
首位两指针：
向内移动短板，移动后板可能变长，则面积可能增大（因为 j-i 变小了，所以不一定会增大，即可能）
			移动后板可能变短或不变，面积一定变小
向内移动长板，移动后板可能变长，但是 min(h[i],h[j]) 不变（水容量取决于短板），面积一定变小
			移动后板可能变短或不变，面积一定变小
总结：移动短板可能变大，移动长板一定变小，所以我们每次都移动短板，直到找到最大面积

https://leetcode-cn.com/problems/container-with-most-water/solution/container-with-most-water-shuang-zhi-zhen-fa-yi-do/
*/

// 双指针法
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i != j {
		tmp := min(height[i], height[j]) * (j - i)
		if tmp > res {
			res = tmp
		}
		// 始终移动短板
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
