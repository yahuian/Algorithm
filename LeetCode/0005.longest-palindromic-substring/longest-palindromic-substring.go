package _0005

// 中心拓展法
// O(n^2)
func longestPalindrome(s string) string {
	if len(s) < 0 || s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := extend(&s, i, i)
		len2 := extend(&s, i, i+1)
		length := max(len1, len2)
		if length > end-start {
			// 此处length-1是为了同意以1个为中心还是以2个为中心
			// 可以分别画一下，就知道了
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}

// 该函数返回以[left,right]为中心的最长回文串的长度
// 当left==right时，可以理解为以1个字符为中心，比如aba
// 当left+1==right时，可以理解为以2个字符为中心，比如abba
// 此处的*string可以减少字符串复制，减少开销
func extend(s *string, left, right int) int {
	for left >= 0 && right < len(*s) && (*s)[left] == (*s)[right] {
		left--
		right++
	}
	// 此处为什么要减一，是因为我们跳出循环的时候，left和right已经都被--了，相当于它们的距离比原来大1
	// 所以我们要在最后的时候-1，恢复为正确的距离
	return right - left - 1
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
