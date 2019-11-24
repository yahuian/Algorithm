package _0089

// 动态规划的思想
// n的结果可以由n-1的结果得到
func grayCode(n int) []int {
	res := []int{0}
	for i := 1; i <= n; i++ {
		e := 1 << (i - 1)
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, e+res[j])
		}
	}
	return res
}

// https://leetcode-cn.com/problems/gray-code/solution/gray-code-jing-xiang-fan-she-fa-by-jyd/
// https://leetcode-cn.com/problems/gray-code/solution/c-dong-tai-gui-hua-jian-ji-yi-dong-shi-jian-ji-kon/
