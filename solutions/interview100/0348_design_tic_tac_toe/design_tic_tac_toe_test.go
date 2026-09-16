package designtictactoe

import "testing"

func TestTicTacToe(t *testing.T) {
	// move 描述一次落子：行、列、玩家
	type move struct {
		row, col, player int
	}

	tests := []struct {
		name  string
		n     int
		moves []move
		want  []int
	}{
		{
			// LeetCode 官方示例：3x3，玩家 1 在第 2 行获胜
			"示例: 3x3玩家1行获胜",
			3,
			[]move{{0, 0, 1}, {0, 2, 2}, {2, 2, 1}, {1, 1, 2}, {2, 0, 1}, {1, 0, 2}, {2, 1, 1}},
			[]int{0, 0, 0, 0, 0, 0, 1},
		},
		{
			// 2x2：玩家 1 占满主对角线获胜
			"2x2玩家1主对角线获胜",
			2,
			[]move{{0, 0, 1}, {0, 1, 2}, {1, 1, 1}},
			[]int{0, 0, 1},
		},
		{
			// 玩家 1 占满一列获胜
			"玩家1列获胜",
			3,
			[]move{{0, 1, 1}, {0, 0, 2}, {1, 1, 1}, {2, 2, 2}, {2, 1, 1}},
			[]int{0, 0, 0, 0, 1},
		},
		{
			// 玩家 2 占满副对角线获胜
			"玩家2副对角线获胜",
			3,
			[]move{{0, 0, 1}, {0, 2, 2}, {1, 0, 1}, {1, 1, 2}, {2, 2, 1}, {2, 0, 2}},
			[]int{0, 0, 0, 0, 0, 2},
		},
		{
			// 边界：最小棋盘 2x2 首步不可能获胜
			"2x2单步无胜局",
			2,
			[]move{{0, 0, 1}},
			[]int{0},
		},
		{
			// 交替落子但互不干扰，直到最后一刻玩家 1 行获胜
			"4x4玩家1末步行获胜",
			4,
			[]move{{3, 0, 1}, {0, 0, 2}, {3, 1, 1}, {0, 1, 2}, {3, 2, 1}, {0, 2, 2}, {3, 3, 1}},
			[]int{0, 0, 0, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := Constructor(tt.n)
			for i, mv := range tt.moves {
				got := game.Move(mv.row, mv.col, mv.player)
				if got != tt.want[i] {
					t.Errorf("第 %d 步 move(%d, %d, %d) = %d, want %d", i, mv.row, mv.col, mv.player, got, tt.want[i])
				}
			}
		})
	}
}

func BenchmarkTicTacToeMove(b *testing.B) {
	// 交替落子填满 100x100 棋盘的场景（不触发获胜即可测满盘性能）
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		game := Constructor(100)
		b.StartTimer()
		for r := 0; r < 100; r++ {
			// 每行先由玩家 1 落子，再由玩家 2 落子，避免一方提前占满
			for c := 0; c < 100; c++ {
				player := 1
				if c%2 == 1 {
					player = 2
				}
				game.Move(r, c, player)
			}
		}
	}
}
