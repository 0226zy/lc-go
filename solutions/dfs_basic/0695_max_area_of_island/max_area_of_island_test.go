package maxareaofisland

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: 最大面积为6",
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			6,
		},
		{
			"示例2: 全为水",
			[][]int{{0, 0, 0, 0, 0, 0, 0, 0}},
			0,
		},

		// 边界情况
		{"空网格", [][]int{}, 0},
		{"单格陆地", [][]int{{1}}, 1},
		{"单格水", [][]int{{0}}, 0},
		{
			"全为陆地",
			[][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			9,
		},
		{
			"多座小岛取最大",
			[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			},
			4,
		},
		{
			"对角不相连",
			[][]int{
				{1, 0},
				{0, 1},
			},
			1,
		},
		{
			"蛇形长岛",
			[][]int{
				{1, 1, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 1, 1},
				{0, 0, 0, 0, 1},
			},
			8,
		},
		{
			"横条岛屿",
			[][]int{{1, 1, 1, 1, 1}},
			5,
		},
		{
			"竖条岛屿",
			[][]int{{1}, {1}, {1}, {1}},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制网格，因为 DFS 会原地修改网格（淹岛）
			grid := make([][]int, len(tt.grid))
			for i := range tt.grid {
				grid[i] = make([]int, len(tt.grid[i]))
				copy(grid[i], tt.grid[i])
			}
			if got := MaxAreaOfIsland(grid); got != tt.want {
				t.Errorf("MaxAreaOfIsland() = %d, want %d", got, tt.want)
			}
		})
	}
}

// makeGrid 生成 rows x cols 的网格，全部填充 fill 值
func makeGrid(rows, cols, fill int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = fill
		}
	}
	return grid
}

func BenchmarkMaxAreaOfIsland(b *testing.B) {
	benchmarks := []struct {
		name string
		rows int
		cols int
		fill int
	}{
		{"10x10全水", 10, 10, 0},
		{"50x50全陆地", 50, 50, 1},
		{"50x50全陆地极限", 50, 50, 1},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxAreaOfIsland(makeGrid(bm.rows, bm.cols, bm.fill))
			}
		})
	}
}
