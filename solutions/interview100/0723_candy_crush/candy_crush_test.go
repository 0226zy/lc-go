package candycrush

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

// copyBoard 深拷贝棋盘，避免 CandyCrush 原地修改影响其他用例
func copyBoard(board [][]int) [][]int {
	cp := make([][]int, len(board))
	for i := range board {
		cp[i] = make([]int, len(board[i]))
		copy(cp[i], board[i])
	}
	return cp
}

func TestCandyCrush(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  [][]int
	}{
		// LeetCode 官方示例 1：大型连锁粉碎
		{
			"示例1: 官方大棋盘",
			[][]int{
				{110, 5, 112, 113, 114},
				{210, 211, 5, 213, 214},
				{310, 311, 3, 313, 314},
				{410, 411, 412, 5, 414},
				{5, 1, 512, 3, 3},
				{610, 4, 1, 613, 614},
				{710, 1, 2, 713, 714},
				{810, 1, 2, 1, 1},
				{1, 1, 2, 2, 2},
				{4, 1, 4, 4, 1014},
			},
			[][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{110, 0, 0, 0, 114},
				{210, 0, 0, 0, 214},
				{310, 0, 0, 113, 314},
				{410, 0, 0, 213, 414},
				{610, 211, 112, 313, 614},
				{710, 311, 412, 613, 714},
				{810, 411, 512, 713, 1014},
			},
		},
		// LeetCode 官方示例 2：下落后引发第二轮粉碎
		{
			"示例2: 下落引发连锁",
			[][]int{
				{1, 3, 5, 5, 2},
				{3, 4, 3, 3, 1},
				{3, 2, 4, 5, 2},
				{2, 4, 4, 5, 5},
				{1, 4, 4, 1, 1},
			},
			[][]int{
				{1, 3, 0, 0, 0},
				{3, 4, 0, 5, 2},
				{3, 2, 0, 3, 1},
				{2, 4, 0, 5, 2},
				{1, 4, 3, 1, 1},
			},
		},
		// 边界：3x3 最小棋盘，无任何可粉碎糖果
		{
			"边界: 3x3 无可粉碎",
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		// 边界：整行 3 个相同糖果，横向粉碎
		{
			"边界: 整行横向三连",
			[][]int{
				{7, 7, 7},
				{1, 2, 3},
				{4, 5, 6},
			},
			[][]int{
				{0, 0, 0},
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		// 边界：整列 3 个相同糖果，纵向粉碎后整列变空
		{
			"边界: 整列纵向三连",
			[][]int{
				{5, 1, 2},
				{5, 3, 4},
				{5, 6, 7},
			},
			[][]int{
				{0, 1, 2},
				{0, 3, 4},
				{0, 6, 7},
			},
		},
		// 边界：连续 4 个相同糖果应一次全部粉碎
		{
			"边界: 横向四连全部粉碎",
			[][]int{
				{9, 9, 9, 9},
				{1, 2, 3, 4},
				{5, 6, 7, 8},
			},
			[][]int{
				{0, 0, 0, 0},
				{1, 2, 3, 4},
				{5, 6, 7, 8},
			},
		},
		// 边界：十字交叉处的糖果横竖同时被标记，但只粉碎一次
		{
			"边界: 十字交叉同时粉碎",
			[][]int{
				{1, 2, 1},
				{2, 2, 2},
				{1, 2, 1},
			},
			[][]int{
				{0, 0, 0},
				{1, 0, 1},
				{1, 0, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CandyCrush(copyBoard(tt.board)); !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("CandyCrush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCandyCrush(b *testing.B) {
	// 构造一个 50x50 的最大规模棋盘（无三连，单次扫描即稳定）
	board := make([][]int, 50)
	for i := range board {
		board[i] = make([]int, 50)
		for j := range board[i] {
			// 保证水平、垂直方向都不出现三连
			board[i][j] = (i*3+j*7)%1999 + 1
		}
	}
	// 校验构造的棋盘确实没有三连，否则基准会包含粉碎逻辑
	probe := copyBoard(board)
	if got := CandyCrush(probe); !utils.Equal2DIntSlice(got, board) {
		b.Skip("基准棋盘构造不稳定，跳过")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CandyCrush(copyBoard(board))
	}
}
