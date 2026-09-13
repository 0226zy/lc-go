package sudokusolver

// SolveSudoku 解数独
// 给定一个 9x9 的数独棋盘，空格用 '.' 表示，编写程序通过填充空格来解数独，
// 要求每行、每列、每个 3x3 宫内的数字均为 1-9 且不重复，原地修改棋盘。
// 时间复杂度: O(9^m)，m 为空格数量，实际搜索因剪枝远小于上界  空间复杂度: O(m) 递归栈深度
func SolveSudoku(board [][]byte) {
	// row[i][d] 表示第 i 行是否已使用数字 d，col 与 box 同理
	var row, col [9][9]bool
	var box [3][3][9]bool

	// 收集所有空格位置，扫描一遍棋盘初始化三个布尔数组
	type point struct{ r, c int }
	empty := []point{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				empty = append(empty, point{i, j})
				continue
			}
			d := board[i][j] - '1'
			row[i][d] = true
			col[j][d] = true
			box[i/3][j/3][d] = true
		}
	}

	var backtrack func(idx int) bool
	backtrack = func(idx int) bool {
		if idx == len(empty) {
			return true // 所有空格都已填满
		}
		p := empty[idx]
		for d := 0; d < 9; d++ {
			// 剪枝：行、列、宫中任一已出现该数字则跳过
			if row[p.r][d] || col[p.c][d] || box[p.r/3][p.c/3][d] {
				continue
			}
			board[p.r][p.c] = byte('1' + d)
			row[p.r][d], col[p.c][d], box[p.r/3][p.c/3][d] = true, true, true
			if backtrack(idx + 1) {
				return true // 找到解，沿递归链一路返回，保留棋盘状态
			}
			board[p.r][p.c] = '.'
			row[p.r][d], col[p.c][d], box[p.r/3][p.c/3][d] = false, false, false
		}
		return false // 当前空格无解，触发回溯
	}
	backtrack(0)
}
