package numberofdistinctislands

import "testing"

// cloneGrid 深拷贝网格，避免实现函数原地修改影响其他用例
func cloneGrid(grid [][]int) [][]int {
	g := make([][]int, len(grid))
	for i := range grid {
		g[i] = make([]int, len(grid[i]))
		copy(g[i], grid[i])
	}
	return g
}

func TestNumDistinctIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 两个相同的2x2方块", [][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 0, 1, 1},
			{0, 0, 0, 1, 1},
		}, 1},
		{"示例2: 三个形状互不相同", [][]int{
			{1, 1, 0, 1, 1},
			{1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1},
			{1, 1, 0, 1, 1},
		}, 3},

		// 边界：空网格与全水
		{"空网格", [][]int{}, 0},
		{"全是水", [][]int{
			{0, 0},
			{0, 0},
		}, 0},

		// 边界：单格
		{"单个陆地", [][]int{{1}}, 1},
		{"单个水", [][]int{{0}}, 0},

		// 边界：全陆地，只有一个大岛屿
		{"全是陆地", [][]int{
			{1, 1, 1},
			{1, 1, 1},
		}, 1},

		// 旋转/翻转不算相同形状
		{"L形与其旋转视为不同", [][]int{
			{1, 0, 0, 1, 1},
			{1, 0, 0, 1, 0},
			{1, 1, 0, 0, 0},
		}, 2},

		// 平移相同算一种
		{"两个一字形平移相同", [][]int{
			{1, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 1, 1, 1},
		}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumDistinctIslands(cloneGrid(tt.grid)); got != tt.want {
				t.Errorf("NumDistinctIslands() = %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkNumDistinctIslands(b *testing.B) {
	// 构造 50x50 棋盘格式网格，包含大量小岛屿
	grid := make([][]int, 50)
	for i := range grid {
		grid[i] = make([]int, 50)
		for j := range grid[i] {
			if (i+j)%2 == 0 {
				grid[i][j] = 1
			}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NumDistinctIslands(cloneGrid(grid))
	}
}
