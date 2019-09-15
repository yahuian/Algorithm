package _0070_

// 递归版，LeetCode会超时
//func climbStairs(n int) int {
//	if n == 1 {
//		return 1
//	}
//	if n == 2 {
//		return 2
//	}
//	return climbStairs(n-1) + climbStairs(n-2)
//}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1 // 爬第1阶有1种方法
	dp[2] = 2 // 爬第2阶有2中方法
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		// 爬第i阶有2种方法，要么从第i-1阶一步爬上来
		// 要么从第i-2阶一次跨2步上来
	}
	return dp[n]
}
