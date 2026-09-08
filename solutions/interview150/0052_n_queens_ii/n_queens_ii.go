package nqueensii

// TotalNQueens N 皇后 II
// 在 n×n 的棋盘上放置 n 个皇后，使它们互不攻击（不同行、不同列、不同对角线），返回方案总数。
// 使用位运算标记列与两条对角线的占用情况。
// 时间复杂度: O(n!)  空间复杂度: O(n) 递归栈深度
func TotalNQueens(n int) int {
	count := 0
	// cols/diag1/diag2 的第 j 位为 1 表示该列/对角线已被占用
	// diag1 方向为 r-c 恒定，diag2 方向为 r+c 恒定
	var backtrack func(row, cols, diag1, diag2 int)
	backtrack = func(row, cols, diag1, diag2 int) {
		if row == n {
			count++
			return
		}
		// 当前行所有可放位置：低 n 位中未被占用的位
		available := ((1 << n) - 1) &^ (cols | diag1 | diag2)
		for available > 0 {
			pos := available & (-available) // 取最低位的可用位置
			available &= available - 1      // 清除该位，尝试下一个
			backtrack(row+1, cols|pos, (diag1|pos)<<1, (diag2|pos)>>1)
		}
	}
	backtrack(0, 0, 0, 0)
	return count
}
