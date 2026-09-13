package islandperimeter

import "testing"

// cloneGrid 复制网格，因为 DFS 主解会染色修改网格
func cloneGrid(grid [][]int) [][]int {
	cloned := make([][]int, len(grid))
	for i := range grid {
		cloned[i] = make([]int, len(grid[i]))
		copy(cloned[i], grid[i])
	}
	return cloned
}

func TestIslandPerimeter(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: 带凸起的岛屿",
			[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			16,
		},
		{
			"示例2: 单格陆地",
			[][]int{{1}},
			4,
		},
		{
			"示例3: 单格陆地带水",
			[][]int{{1, 0}},
			4,
		},

		// 边界情况
		{"空网格", [][]int{}, 0},
		{"全为水域", [][]int{{0, 0}, {0, 0}}, 0},
		{
			"2x2全陆地",
			[][]int{
				{1, 1},
				{1, 1},
			},
			8,
		},
		{
			"横向长条",
			[][]int{{1, 1, 1, 1, 1}},
			12, // 2*(5+1)
		},
		{
			"竖向长条",
			[][]int{{1}, {1}, {1}},
			8, // 2*(3+1)
		},
		{
			"大块正方形",
			[][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			12, // 4*3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IslandPerimeter(cloneGrid(tt.grid)); got != tt.want {
				t.Errorf("IslandPerimeter() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestIslandPerimeterCount(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			"示例1: 带凸起的岛屿",
			[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			16,
		},
		{"示例2: 单格陆地", [][]int{{1}}, 4},
		{"示例3: 单格陆地带水", [][]int{{1, 0}}, 4},
		{"空网格", [][]int{}, 0},
		{"全为水域", [][]int{{0, 0}, {0, 0}}, 0},
		{
			"2x2全陆地",
			[][]int{
				{1, 1},
				{1, 1},
			},
			8,
		},
		{"横向长条", [][]int{{1, 1, 1, 1, 1}}, 12},
		{
			"大块正方形",
			[][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IslandPerimeterCount(tt.grid); got != tt.want {
				t.Errorf("IslandPerimeterCount() = %d, want %d", got, tt.want)
			}
		})
	}
}

// makeLandGrid 生成 rows x cols 的全陆地网格
func makeLandGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = 1
		}
	}
	return grid
}

func BenchmarkIslandPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IslandPerimeter(makeLandGrid(100, 100))
	}
}

func BenchmarkIslandPerimeterCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IslandPerimeterCount(makeLandGrid(100, 100))
	}
}
