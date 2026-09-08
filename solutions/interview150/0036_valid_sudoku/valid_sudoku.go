package validsudoku

// IsValidSudoku 有效的数独
// 判断 9x9 的数独盘面是否有效：每一行、每一列、每一个 3x3 宫内都不出现重复数字。
// 盘面中的空位用 '.' 表示，填入的数字为 '1'..'9'。
// 时间复杂度: O(9x9)=O(1) 常数遍历  空间复杂度: O(9x3)=O(1) 常数辅助空间
func IsValidSudoku(board [][]byte) bool {
	// rows[i] 第 i 行出现过的数字集合，用位掩码表示
	// cols[j] 第 j 列出现过的数字集合
	// boxes[k] 第 k 个 3x3 宫出现过的数字集合
	var rows, cols, boxes [9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			// 数字 '1'..'9' 映射到第 0..8 位
			bit := 1 << (board[i][j] - '1')
			// 宫索引：行号按 3 行分组，列号按 3 列分组
			box := (i/3)*3 + j/3
			// 任意一行、一列、一宫中出现过该数字即为非法
			if rows[i]&bit != 0 || cols[j]&bit != 0 || boxes[box]&bit != 0 {
				return false
			}
			rows[i] |= bit
			cols[j] |= bit
			boxes[box] |= bit
		}
	}
	return true
}
