package _0120

// 递归版本,leetcode会超时
//func minimumTotal(triangle [][]int) int {
//	return travel(triangle, 0, 0)
//}
//
//func travel(triangle [][]int, i, j int) int {
//	if i == len(triangle) {
//		return 0
//	}
//	return triangle[i][j] + min(travel(triangle, i+1, j), travel(triangle, i+1, j+1))
//}
//

// 动态规划
// 状态的定义：dp[i,j]表示自最底层向上走，到达点triangle[i,j]中的所有路径中，值最小的那条路径的长度
// 状态转移方程：dp[i,j] = triangle[i][j] + min(dp[i+1,j], dp[i+1,j+1])
// 初始状态：dp[rowMax,j] = triangle[rowMax,j]
func minimumTotal(triangle [][]int) int {
	rowMax := len(triangle)
	// 从倒数第二层开始
	for i := rowMax - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// 直接在原数组上修改，没有单独开辟一个dp [][]int来存节点状态
			triangle[i][j] = triangle[i][j] + min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
