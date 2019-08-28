package _0239

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || k <= 0 {
		return []int{}
	}
	var window, result []int
	for index, value := range nums {
		// 检查每个数在window中的生命周期是否合法
		if index >= k && window[0] <= index-k {
			window = window[1:]
		}
		for len(window) != 0 && value > nums[window[len(window)-1]] {
			window = window[:len(window)-1]
		}
		// 注意window双端队列里面存放的是nums的索引
		window = append(window, index)
		if index >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}

// 使用双端队列
