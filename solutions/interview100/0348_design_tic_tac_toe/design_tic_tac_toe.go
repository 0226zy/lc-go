package designtictactoe

// TicTacToe 设计井字棋
// 在 n x n 棋盘上两名玩家轮流落子，某玩家在任意一行、一列或一条对角线上占满 n 格即获胜。
// 时间复杂度: 每次 move O(1)  空间复杂度: O(n)
type TicTacToe struct {
	n        int
	rows     []int // 每行的净计数：玩家1 +1，玩家2 -1
	cols     []int // 每列的净计数
	diag     int   // 主对角线（row == col）净计数
	antiDiag int   // 副对角线（row + col == n-1）净计数
}

// Constructor 初始化一个 n x n 的井字棋棋盘
func Constructor(n int) TicTacToe {
	return TicTacToe{
		n:    n,
		rows: make([]int, n),
		cols: make([]int, n),
	}
}

// Move 玩家 player（1 或 2）在 (row, col) 落子
// 返回值：0 表示无人获胜，1 表示玩家 1 获胜，2 表示玩家 2 获胜
func (t *TicTacToe) Move(row int, col int, player int) int {
	// 玩家 1 贡献 +1，玩家 2 贡献 -1；某条线计数绝对值达到 n 即被同一玩家占满
	delta := 1
	if player == 2 {
		delta = -1
	}
	t.rows[row] += delta
	t.cols[col] += delta
	if row == col {
		t.diag += delta
	}
	if row+col == t.n-1 {
		t.antiDiag += delta
	}
	if abs(t.rows[row]) == t.n || abs(t.cols[col]) == t.n ||
		abs(t.diag) == t.n || abs(t.antiDiag) == t.n {
		return player
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
