package wordsearch

// dirs 四个移动方向：上、下、左、右
var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// Exist 单词搜索
// 给定一个 m x n 的字符网格 board 和一个字符串单词 word，
// 如果 word 存在于网格中（相邻格子横纵相连、同一格子只能用一次），返回 true。
// 时间复杂度: O(m * n * 4^L)  L 为单词长度  空间复杂度: O(L) 递归栈深度
func Exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var backtrack func(i, j, k int) bool
	backtrack = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true // 最后一个字符也匹配，找到单词
		}
		// 标记为已访问（0 不会出现在棋盘里），避免同一路径重复使用
		temp := board[i][j]
		board[i][j] = 0
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && backtrack(ni, nj, k+1) {
				board[i][j] = temp // 还原现场
				return true
			}
		}
		board[i][j] = temp // 回溯：恢复原字符
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}
