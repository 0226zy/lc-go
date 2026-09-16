package shortestdistancefromallbuildings

import "testing"

func TestShortestDistance(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: 3x5 网格三栋建筑物",
			[][]int{
				{1, 0, 2, 0, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0},
			},
			7,
		},
		{
			"示例2: 一栋建筑物一块空地",
			[][]int{{1, 0}},
			1,
		},
		{
			"示例3: 没有空地",
			[][]int{{1}},
			-1,
		},

		// 边界：建筑物被障碍物完全包围，任何空地都无法到达它
		{
			"建筑物被障碍物完全隔离",
			[][]int{
				{2, 2, 2},
				{2, 1, 2},
				{2, 2, 2},
				{0, 0, 0},
			},
			-1,
		},

		// 边界：空地彼此不连通，各自只能到达部分建筑物
		{
			"空地无法到达全部建筑物",
			[][]int{
				{1, 2, 1},
				{0, 2, 0},
			},
			-1,
		},

		// 典型场景：两栋建筑物，最佳位置在中间
		{
			"两栋建筑物夹一块空地",
			[][]int{{1, 0, 1}},
			2,
		},

		// 典型场景：十字形布局，中心最优
		{
			"四栋建筑物环绕中心",
			[][]int{
				{0, 1, 0},
				{1, 0, 1},
				{0, 1, 0},
			},
			4,
		},

		// 边界：只有建筑物和障碍物，无空地
		{
			"全是建筑物与障碍物",
			[][]int{
				{1, 2},
				{2, 1},
			},
			-1,
		},

		// 典型场景：1x4 一排
		{
			"一排中的最优位置",
			[][]int{{1, 0, 0, 1}},
			3, // 位置 1 或 2，距离之和均为 1+2=3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortestDistance(tt.grid); got != tt.want {
				t.Errorf("ShortestDistance(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func BenchmarkShortestDistance(b *testing.B) {
	// 构造 50x50 网格：棋盘状分布建筑物与空地
	size := 50
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if (i+j)%10 == 0 {
				grid[i][j] = 1
			}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShortestDistance(grid)
	}
}
