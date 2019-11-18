package _0046

func permute(nums []int) [][]int {
	n := len(nums)
	vector := make([]int, n)
	used := make([]bool, n)
	var ans [][]int

	makePermute(0, n, vector, nums, used, &ans)

	return ans
}

func makePermute(cur, n int, vector, nums []int, used []bool, ans *[][]int) {
	if cur == n {
		temp := make([]int, n)
		copy(temp, vector)
		*ans = append(*ans, temp)
		return
	}

	for i := 0; i < n; i++ {
		if !used[i] {
			used[i] = true
			vector[cur] = nums[i]
			makePermute(cur+1, n, vector, nums, used, ans)
			used[i] = false
		}
	}
}

// 参考资料：
// https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0046.permutations/permutations.go
// https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
