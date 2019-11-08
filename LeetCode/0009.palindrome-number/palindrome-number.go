package _0009

func isPalindrome(x int) bool {
	// 负数一定不是回文数
	if x < 0 {
		return false
	}
	// 如果x最后一位是0，则第一位也要是0，只有0才符合
	if x%10 == 0 && x != 0 {
		return false
	}
	reverseNumber := 0
	// 当 x > reverseNumber时，表示还没有反转一半
	for x > reverseNumber {
		pop := x % 10
		reverseNumber = reverseNumber*10 + pop
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	// https://leetcode-cn.com/problems/palindrome-number/solution/hui-wen-shu-by-leetcode/
	return x == reverseNumber || x == reverseNumber/10
}
