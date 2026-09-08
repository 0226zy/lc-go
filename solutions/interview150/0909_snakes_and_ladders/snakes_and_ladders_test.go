package snakesandladders

import "testing"

func TestSnakesAndLadders(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  int
	}{
		// LeetCode 官方示例
		{"示例1: 6x6梯子", [][]int{
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 35, -1, -1, 13, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 15, -1, -1, -1, -1},
		}, 4},
		{"示例2: 2x2梯子", [][]int{{-1, -1}, {-1, 3}}, 1},

		// 边界：最小棋盘无蛇梯
		{"2x2全空", [][]int{{-1, -1}, {-1, -1}}, 1},

		// 蛇梯送回起点导致无法到达
		{"梯子送回起点", [][]int{
			{1, 1, -1},
			{1, 1, 1},
			{-1, 1, 1},
		}, -1},

		// 蛇梯直达终点：1 -> 7 -> 8(梯子到16)
		{"连续梯子", [][]int{
			{-1, -1, -1, -1},
			{-1, -1, -1, -1},
			{16, -1, -1, -1},
			{-1, 8, -1, -1},
		}, 2},

		// 棋盘较大且无蛇梯：纯掷骰子
		{"5x5全空", [][]int{
			{-1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1},
		}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakesAndLadders(tt.board); got != tt.want {
				t.Errorf("SnakesAndLadders() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestSquareToCoord(t *testing.T) {
	tests := []struct {
		name   string
		pos, n int
		wantR  int
		wantC  int
	}{
		{"6x6编号1", 1, 6, 5, 0},
		{"6x6编号2", 2, 6, 5, 1},
		{"6x6编号6", 6, 6, 5, 5},
		{"6x6编号7(蛇形折返)", 7, 6, 4, 5},
		{"6x6编号12", 12, 6, 4, 0},
		{"6x6编号13", 13, 6, 3, 0},
		{"6x6编号14", 14, 6, 3, 1},
		{"6x6编号36", 36, 6, 0, 0},
		{"2x2编号1", 1, 2, 1, 0},
		{"2x2编号2", 2, 2, 1, 1},
		{"2x2编号3", 3, 2, 0, 1},
		{"2x2编号4", 4, 2, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, c := squareToCoord(tt.pos, tt.n)
			if r != tt.wantR || c != tt.wantC {
				t.Errorf("squareToCoord(%d, %d) = (%d, %d), want (%d, %d)",
					tt.pos, tt.n, r, c, tt.wantR, tt.wantC)
			}
		})
	}
}

func BenchmarkSnakesAndLadders(b *testing.B) {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SnakesAndLadders(board)
	}
}
