package candycrush

// CandyCrush 粉碎糖果
// 给定 m x n 的棋盘 board，不断将水平或垂直方向上连续 3 个及以上相同类型的糖果同时粉碎，
// 并让上方糖果受重力下落填补空位，重复该过程直到棋盘稳定，返回最终棋盘。
// 时间复杂度: O(k*m*n)，k 为粉碎轮数（每轮至少消除 3 个糖果）  空间复杂度: O(m*n)，标记数组
func CandyCrush(board [][]int) [][]int {
	m, n := len(board), len(board[0])
	for {
		// 标记阶段：找出本轮所有需要粉碎的格子
		marked := make([][]bool, m)
		for i := range marked {
			marked[i] = make([]bool, n)
		}
		found := false

		// 横向：检查每个长度为 3 的窗口，连续更长的情况会被相邻窗口覆盖
		for i := 0; i < m; i++ {
			for j := 0; j+2 < n; j++ {
				v := board[i][j]
				if v != 0 && v == board[i][j+1] && v == board[i][j+2] {
					marked[i][j], marked[i][j+1], marked[i][j+2] = true, true, true
					found = true
				}
			}
		}
		// 纵向：同理检查每列长度为 3 的窗口
		for j := 0; j < n; j++ {
			for i := 0; i+2 < m; i++ {
				v := board[i][j]
				if v != 0 && v == board[i+1][j] && v == board[i+2][j] {
					marked[i][j], marked[i+1][j], marked[i+2][j] = true, true, true
					found = true
				}
			}
		}

		// 没有任何可粉碎的糖果，棋盘已稳定
		if !found {
			break
		}

		// 粉碎 + 下落：逐列把未标记的糖果压到列底，顶部填 0
		for j := 0; j < n; j++ {
			w := m - 1 // 写指针，从列底向上
			for i := m - 1; i >= 0; i-- {
				if !marked[i][j] {
					board[w][j] = board[i][j]
					w--
				}
			}
			for ; w >= 0; w-- {
				board[w][j] = 0
			}
		}
	}
	return board
}
