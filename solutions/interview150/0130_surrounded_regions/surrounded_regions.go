package surroundedregions

// Solve 被围绕的区域
// 给你一个 m x n 的矩阵 board，由若干字符 'X' 和 'O' 组成。找到所有被 'X' 围绕的区域，
// 并将这些区域里所有的 'O' 用 'X' 填充。任何边界上的 'O' 都不会被填充为 'X'，
// 任何不在边界上、也不与边界上的 'O' 相连的 'O' 都会被填充为 'X'。
// 时间复杂度: O(m*n) 每个格子最多访问一次  空间复杂度: O(m*n) 递归栈深度最坏情况
func Solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	// 标记与边界相连的 'O' 为 '#'（表示“安全，不能翻转”）
	var mark func(i, j int)
	mark = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = '#'
		mark(i+1, j)
		mark(i-1, j)
		mark(i, j+1)
		mark(i, j-1)
	}

	// 从四条边界出发，把所有与边界相连的 'O' 标记为 '#'
	for i := 0; i < m; i++ {
		mark(i, 0)
		mark(i, n-1)
	}
	for j := 0; j < n; j++ {
		mark(0, j)
		mark(m-1, j)
	}

	// 遍历整个矩阵：'#' 恢复为 'O'，剩下的 'O'（被围绕的）翻转为 'X'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case '#':
				board[i][j] = 'O'
			case 'O':
				board[i][j] = 'X'
			}
		}
	}
}
