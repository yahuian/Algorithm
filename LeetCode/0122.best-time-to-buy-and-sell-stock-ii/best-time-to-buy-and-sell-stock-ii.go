package _0122

// 贪心算法
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var sum int
	for i, j := 0, 1; j < len(prices); i, j = i+1, j+1 {
		if prices[j] > prices[i] {
			sum += prices[j] - prices[i]
		}
	}
	return sum
}
