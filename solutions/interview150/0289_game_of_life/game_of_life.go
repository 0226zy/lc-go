package gameoflife

// GameOfLife 生命游戏
// 原地模拟细胞下一代状态。用复合状态避免修改影响后续统计：
// 2 表示细胞由活变死（1 -> 0），-1 表示细胞由死变活（0 -> 1），最后统一还原。
// 时间复杂度: O(m*n*8) 每个细胞检查 8 个邻居  空间复杂度: O(1)
func GameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	// 8 个邻居的相对偏移
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			live := 0
			for _, d := range dirs {
				x, y := i+d[0], j+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && (board[x][y] == 1 || board[x][y] == 2) {
					live++
				}
			}
			// 活细胞：邻居少于 2 个或超过 3 个则死亡
			if board[i][j] == 1 && (live < 2 || live > 3) {
				board[i][j] = 2
			} else if board[i][j] == 0 && live == 3 {
				// 死细胞：恰好 3 个活邻居则复活
				board[i][j] = -1
			}
		}
	}

	// 统一还原复合状态
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
}

// GameOfLifeCopy 生命游戏（辅助数组解法，供对比）
// 复制一份原矩阵，统计时读原矩阵、写新状态。
// 时间复杂度: O(m*n*8)  空间复杂度: O(m*n) 需要一个同样大小的辅助矩阵
func GameOfLifeCopy(board [][]int) {
	m, n := len(board), len(board[0])
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	copyBoard := make([][]int, m)
	for i := 0; i < m; i++ {
		copyBoard[i] = make([]int, n)
		copy(copyBoard[i], board[i])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			live := 0
			for _, d := range dirs {
				x, y := i+d[0], j+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && copyBoard[x][y] == 1 {
					live++
				}
			}
			if copyBoard[i][j] == 1 {
				if live < 2 || live > 3 {
					board[i][j] = 0
				}
			} else if live == 3 {
				board[i][j] = 1
			}
		}
	}
}
