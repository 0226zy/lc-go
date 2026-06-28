package nqueens

// SolveNQueens N 皇后
// 返回所有不同的 n 皇后问题的解决方案。
// 时间复杂度: O(n!)  空间复杂度: O(n²)
func SolveNQueens(n int) [][]string {
	if n < 1 {
		return nil
	}

	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	ret := make([][]string, 0)
	backtrace(board, 0, &ret)
	return ret
}

func backtrace(board [][]string, row int, ret *[][]string) {
	n := len(board)
	if row == n {
		snapshot := make([]string, n)
		for i := 0; i < n; i++ {
			rowBytes := make([]byte, n)
			for j := 0; j < n; j++ {
				rowBytes[j] = '.'
				if board[i][j] == "Q" {
					rowBytes[j] = 'Q'
				}
			}
			snapshot[i] = string(rowBytes)
		}
		*ret = append(*ret, snapshot)
		return
	}
	for col := 0; col < n; col++ {
		if !isValid(board, row, col) {
			continue
		}
		board[row][col] = "Q"
		backtrace(board, row+1, ret)
		board[row][col] = "."
	}
}

func isValid(board [][]string, row, col int) bool {
	n := len(board)

	// 检查列冲突
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	// 检查左上对角线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	// 检查右上对角线
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
