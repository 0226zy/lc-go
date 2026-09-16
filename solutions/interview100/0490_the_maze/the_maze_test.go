package themaze

import "testing"

// 官方示例 1、2 共用的迷宫
var mazeA = [][]int{
	{0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0},
	{1, 1, 0, 1, 1},
	{0, 0, 0, 0, 0},
}

// 官方示例 3 的迷宫
var mazeB = [][]int{
	{0, 0, 0, 0, 0},
	{1, 1, 0, 0, 1},
	{0, 0, 0, 0, 0},
	{0, 1, 0, 0, 1},
	{0, 1, 0, 0, 0},
}

func TestHasPath(t *testing.T) {
	tests := []struct {
		name        string
		maze        [][]int
		start       []int
		destination []int
		want        bool
	}{
		// LeetCode 官方示例
		{"示例1: 可到达", mazeA, []int{0, 4}, []int{4, 4}, true},
		{"示例2: 无法恰好停下", mazeA, []int{0, 4}, []int{3, 2}, false},
		{"示例3: 目的地不可达", mazeB, []int{4, 3}, []int{0, 1}, false},

		// 边界：单行 / 单列迷宫
		{"单行可滚到端点", [][]int{{0, 0, 0}}, []int{0, 0}, []int{0, 2}, true},
		{"单行只能停在端点", [][]int{{0, 0, 0}}, []int{0, 0}, []int{0, 1}, false},
		{"单列可滚到底", [][]int{{0}, {0}, {0}}, []int{0, 0}, []int{2, 0}, true},

		// 典型场景
		{"被墙壁包围寸步难行", [][]int{{0, 1}, {1, 0}}, []int{0, 0}, []int{1, 1}, false},
		{"途经不算到达", [][]int{{0, 0, 0, 0}}, []int{0, 0}, []int{0, 2}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPath(tt.maze, tt.start, tt.destination); got != tt.want {
				t.Errorf("HasPath(start=%v, destination=%v) = %v, want %v",
					tt.start, tt.destination, got, tt.want)
			}
		})
	}
}

func BenchmarkHasPath(b *testing.B) {
	// 构造 50x50 的空迷宫（无内部墙壁）
	maze := make([][]int, 50)
	for i := range maze {
		maze[i] = make([]int, 50)
	}

	benchmarks := []struct {
		name        string
		maze        [][]int
		start       []int
		destination []int
	}{
		{"5x5官方示例", mazeA, []int{0, 4}, []int{4, 4}},
		{"50x50空迷宫", maze, []int{0, 0}, []int{49, 49}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HasPath(bm.maze, bm.start, bm.destination)
			}
		})
	}
}
