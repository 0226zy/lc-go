package nqueens

// SolveNQueens N 皇后
// 将 n 个皇后放置在 n×n 的棋盘上，使皇后彼此不能互相攻击，返回所有不同的放置方案。
// 时间复杂度: O(n!)  空间复杂度: O(n)
func SolveNQueens(n int) [][]string {
	queens := make([]int, n) // queens[r] 表示第 r 行皇后所在的列
	for i := range queens {
		queens[i] = -1
	}
	cols := make(map[int]bool)  // 已被占用的列
	diag1 := make(map[int]bool) // 主对角线方向：r - c 恒定
	diag2 := make(map[int]bool) // 副对角线方向：r + c 恒定
	result := make([][]string, 0)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			board := make([]string, n)
			for r := 0; r < n; r++ {
				line := make([]byte, n)
				for c := 0; c < n; c++ {
					line[c] = '.'
				}
				line[queens[r]] = 'Q'
				board[r] = string(line)
			}
			result = append(result, board)
			return
		}
		for col := 0; col < n; col++ {
			d1, d2 := row-col, row+col
			if cols[col] || diag1[d1] || diag2[d2] {
				continue // 列或对角线冲突，剪枝
			}
			queens[row] = col
			cols[col], diag1[d1], diag2[d2] = true, true, true
			backtrack(row + 1)
			queens[row] = -1
			delete(cols, col)
			delete(diag1, d1)
			delete(diag2, d2)
		}
	}
	backtrack(0)
	return result
}
