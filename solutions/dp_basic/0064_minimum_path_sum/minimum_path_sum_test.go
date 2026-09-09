package minimumpathsum

import "testing"

var minPathSumCases = []struct {
	name string
	grid [][]int
	want int
}{
	{name: "示例1：3x3", grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, want: 7},
	{name: "示例2：2x3", grid: [][]int{{1, 2, 3}, {4, 5, 6}}, want: 12},
	{name: "单格", grid: [][]int{{1}}, want: 1},
	{name: "2x2走右上更优", grid: [][]int{{1, 2}, {3, 4}}, want: 7},
	{name: "单行只能向右", grid: [][]int{{1, 2, 3}}, want: 6},
	{name: "单列只能向下", grid: [][]int{{1}, {2}, {3}}, want: 6},
	{name: "绕开大数", grid: [][]int{{1, 4, 8}, {6, 2, 7}}, want: 14},
	{name: "含0格子", grid: [][]int{{0, 0}, {0, 0}}, want: 0},
}

func TestMinPathSum(t *testing.T) {
	for _, tt := range minPathSumCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPathSum(tt.grid); got != tt.want {
				t.Errorf("MinPathSum(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestMinPathSumOptimized(t *testing.T) {
	for _, tt := range minPathSumCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPathSumOptimized(tt.grid); got != tt.want {
				t.Errorf("MinPathSumOptimized(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

// 构造一个 50x50 网格用于基准测试
func benchGrid() [][]int {
	grid := make([][]int, 50)
	for i := range grid {
		grid[i] = make([]int, 50)
		for j := range grid[i] {
			grid[i][j] = (i*50 + j) % 10
		}
	}
	return grid
}

func BenchmarkMinPathSum(b *testing.B) {
	grid := benchGrid()
	for i := 0; i < b.N; i++ {
		MinPathSum(grid)
	}
}

func BenchmarkMinPathSumOptimized(b *testing.B) {
	grid := benchGrid()
	for i := 0; i < b.N; i++ {
		MinPathSumOptimized(grid)
	}
}
