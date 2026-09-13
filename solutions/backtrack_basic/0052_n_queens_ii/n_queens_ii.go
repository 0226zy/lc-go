package nqueensii

// TotalNQueens N 皇后 II
// 在 n×n 的棋盘上放置 n 个皇后，使它们互不攻击（不同行、不同列、不同对角线），返回不同方案的总数。
// 使用三个布尔数组分别标记列与两条对角线的占用情况，逐行回溯放置。
// 时间复杂度: O(n!)  空间复杂度: O(n)
func TotalNQueens(n int) int {
	// col[c] 表示第 c 列是否已放置皇后
	col := make([]bool, n)
	// diag1[r-c+n-1] 表示左上→右下方向（r-c 恒定）的对角线是否被占用
	diag1 := make([]bool, 2*n-1)
	// diag2[r+c] 表示右上→左下方向（r+c 恒定）的对角线是否被占用
	diag2 := make([]bool, 2*n-1)

	count := 0
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			count++
			return
		}
		for c := 0; c < n; c++ {
			d1 := row - c + n - 1
			d2 := row + c
			if col[c] || diag1[d1] || diag2[d2] {
				continue
			}
			// 做选择：标记列与两条对角线
			col[c], diag1[d1], diag2[d2] = true, true, true
			backtrack(row + 1)
			// 撤销选择：还原标记
			col[c], diag1[d1], diag2[d2] = false, false, false
		}
	}
	backtrack(0)
	return count
}
