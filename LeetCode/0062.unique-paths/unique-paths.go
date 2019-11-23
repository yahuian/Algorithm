package _0062

// 动态规划
// 状态的定义：dp[i][j]表示从matrix[i][j]位置到终点有多少种路径
// 状态转移方程：dp[i][j]=dp[i][j+1]+dp[i+1][j]
// 初始状态：
// down := len(matrix)-1
// right := len(matrix[0])-1
// dp[down][j] = 1
// dp[i][right] = 1

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 自低向上
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j+1] + dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}
