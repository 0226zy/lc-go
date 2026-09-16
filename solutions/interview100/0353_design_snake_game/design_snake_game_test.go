package designsnakegame

import "testing"

func TestSnakeGame(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
		food   [][]int
		moves  []string
		want   []int
	}{
		{
			// LeetCode 官方示例
			"示例: 3x2棋盘两份食物",
			3, 2, [][]int{{1, 2}, {0, 1}},
			[]string{"R", "D", "R", "U", "L", "U"},
			[]int{0, 0, 1, 1, 2, -1},
		},
		{
			// 咬到自己：吃到长度 4 后头部撞向蛇身中间
			"咬到自己",
			3, 3, [][]int{{0, 1}, {0, 2}, {1, 2}},
			[]string{"R", "R", "D", "L", "R"},
			[]int{1, 2, 3, 3, -1},
		},
		{
			// 边界：第一步就越界
			"第一步向上越界",
			3, 3, nil,
			[]string{"U"},
			[]int{-1},
		},
		{
			// 边界：游戏结束后再移动仍返回 -1
			"游戏结束后继续移动",
			2, 2, nil,
			[]string{"U", "D", "R"},
			[]int{-1, -1, -1},
		},
		{
			// 跟着尾巴走不判撞身：蛇长 2，尾巴让出后头部到达原尾位置
			"跟着尾巴走",
			3, 2, [][]int{{0, 1}},
			[]string{"R", "D", "L", "U", "R", "D"},
			[]int{1, 1, 1, 1, 1, 1},
		},
		{
			// 食物只能按顺序吃：蛇经过后面的食物位置不算吃到
			"食物按顺序出现",
			4, 1, [][]int{{0, 3}, {0, 1}},
			[]string{"R", "R", "R"},
			[]int{0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := Constructor(tt.width, tt.height, tt.food)
			for i, mv := range tt.moves {
				got := game.Move(mv)
				if got != tt.want[i] {
					t.Errorf("第 %d 步 move(%q) = %d, want %d", i, mv, got, tt.want[i])
				}
			}
		})
	}
}

func BenchmarkSnakeGameMove(b *testing.B) {
	// 100x100 棋盘无食物，蛇沿蛇形路径往复移动
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		game := Constructor(100, 100, nil)
		b.StartTimer()
		for row := 0; row < 100; row++ {
			for col := 0; col < 99; col++ {
				dir := "R"
				if row%2 == 1 {
					dir = "L"
				}
				game.Move(dir)
			}
			game.Move("D")
		}
	}
}
