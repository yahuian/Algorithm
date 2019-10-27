package _0097

func exist(board [][]byte, word string) bool {
	row := len(board)
	if row == 0 || len(word) == 0 {
		return false
	}
	column := len(board[0])
	if column == 0 {
		return false
	}

	var dfs func(int, int, int) bool // 函数类型的变量
	dfs = func(r, c, index int) bool {
		if len(word) == index {
			return true
		}
		if r < 0 || c < 0 || r >= row || c >= column || board[r][c] != word[index] {
			return false
		}
		//			       [r-1][c]
		//
		//		[r][c-1]    [r][c]	  [r][c+1]
		//
		//			       [r+1][c]
		//
		// 按照"上"->"右"->"下"->"左"的顺序进行深度优先周游
		// 如果某条路不通则回溯
		tmp := board[r][c]
		board[r][c] = 0 // 标记一下，确保走过的节点不被重复走
		if dfs(r-1, c, index+1) || dfs(r, c+1, index+1) || dfs(r+1, c, index+1) || dfs(r, c-1, index+1) {
			return true
		}
		board[r][c] = tmp
		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0079.word-search/word-search.go
// https://leetcode-cn.com/problems/word-search/solution/zai-er-wei-ping-mian-shang-shi-yong-hui-su-fa-pyth/
