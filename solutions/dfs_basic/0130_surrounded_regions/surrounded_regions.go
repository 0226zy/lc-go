package surroundedregions

// Solve 被围绕的区域
// 给你一个 m x n 的矩阵 board，由 'X' 和 'O' 组成，把所有被 'X' 完全围绕、
// 且不与边界连通的 'O' 区域原地翻转为 'X'。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n) 递归栈最坏深度
func Solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	// mark 从 (i, j) 出发，把与边界连通的 'O' 全部标记为 '#'（安全区）
	var mark func(i, j int)
	mark = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}
		// 先标记再递归，避免重复访问
		board[i][j] = '#'
		mark(i-1, j)
		mark(i+1, j)
		mark(i, j-1)
		mark(i, j+1)
	}

	// 从四条边界上的每个 'O' 出发反向标记
	for i := 0; i < m; i++ {
		mark(i, 0)
		mark(i, n-1)
	}
	for j := 0; j < n; j++ {
		mark(0, j)
		mark(m-1, j)
	}

	// 收尾遍历：安全区 '#' 还原为 'O'，被围绕的 'O' 翻转为 'X'
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
