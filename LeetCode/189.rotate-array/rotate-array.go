package _189

func rotate(nums []int, k int) {
	// 输入参数合法性校验
	if k < 0 {
		return
	}
	if len(nums) < k {
		k = k - len(nums)
	}

	// 将一个数组右移k位的最简单方法是连续调用reverse函数三次
	// reverse(s)
	// reverse[:k]
	// reverse[k:]
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

	// 将一个数组左移k位
	// reverse(s[:k])
	// reverse(s[k:])
	// reverse(s)
}

// 就地反转一个slice中的元素
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
