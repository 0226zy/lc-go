package themazeii

import "testing"

func TestShortestDistance(t *testing.T) {
	maze1 := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	maze3 := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}

	tests := []struct {
		name        string
		maze        [][]int
		start       []int
		destination []int
		want        int
	}{
		// LeetCode 官方示例
		{"示例1: 最短距离12", maze1, []int{0, 4}, []int{4, 4}, 12},
		{"示例2: 无法停下", maze1, []int{0, 4}, []int{3, 2}, -1},
		{"示例3: 不可达", maze3, []int{4, 3}, []int{0, 1}, -1},

		// 边界：起点即终点（不滚动，距离为 0）
		{"起点即终点", [][]int{{0, 0}, {0, 0}}, []int{0, 0}, []int{0, 0}, 0},

		// 边界：单行迷宫
		{"单行直达", [][]int{{0, 0, 0}}, []int{0, 0}, []int{0, 2}, 2},
		{"单行被墙隔开", [][]int{{0, 1, 0}}, []int{0, 0}, []int{0, 2}, -1},

		// 边界：单列迷宫
		{"单列直达", [][]int{{0}, {0}, {0}}, []int{0, 0}, []int{2, 0}, 2},

		// 典型场景：滚过头停不下来（目标格不是墙边）
		{"中间格停不住", [][]int{{0, 0, 0}}, []int{0, 0}, []int{0, 1}, -1},

		// 典型场景：需要折返多次滚动
		// 路径：(0,0)→右滚2步→(0,2)→下滚2步→(2,2)→左滚2步→(2,0)，共 6 步
		{"需要绕路", [][]int{
			{0, 0, 0},
			{1, 1, 0},
			{0, 0, 0},
		}, []int{0, 0}, []int{2, 0}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortestDistance(tt.maze, tt.start, tt.destination); got != tt.want {
				t.Errorf("ShortestDistance(%v, %v, %v) = %d, want %d",
					tt.maze, tt.start, tt.destination, got, tt.want)
			}
		})
	}
}

func BenchmarkShortestDistance(b *testing.B) {
	// 构造 50×50 的空旷迷宫（无墙）
	maze := make([][]int, 50)
	for i := range maze {
		maze[i] = make([]int, 50)
	}
	start := []int{0, 0}
	destination := []int{49, 49}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShortestDistance(maze, start, destination)
	}
}
