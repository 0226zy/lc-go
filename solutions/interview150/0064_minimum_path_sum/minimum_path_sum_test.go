package minimumpathsum

import "testing"

var minPathSumCases = []struct {
	name string
	grid [][]int
	want int
}{
	// LeetCode 官方示例
	{name: "示例1: 3x3网格", grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, want: 7},
	{name: "示例2: 2x3网格", grid: [][]int{{1, 2, 3}, {4, 5, 6}}, want: 12},

	// 边界：单元素与单行/单列
	{name: "单元素网格", grid: [][]int{{5}}, want: 5},
	{name: "单行网格只能向右", grid: [][]int{{1, 2, 3, 4}}, want: 10},
	{name: "单列网格只能向下", grid: [][]int{{1}, {2}, {3}, {4}}, want: 10},

	// 边界：全零网格
	{name: "全零2x2网格", grid: [][]int{{0, 0}, {0, 0}}, want: 0},
	{name: "起点为0", grid: [][]int{{0, 0}, {1, 1}}, want: 1},

	// 普通场景：绕开高代价区域
	{name: "2x2走右上更优", grid: [][]int{{1, 2}, {3, 4}}, want: 7},
	{name: "绕开中间大数", grid: [][]int{{1, 100, 1}, {1, 100, 1}, {1, 1, 1}}, want: 5},
	{name: "先下后右更优", grid: [][]int{{1, 9, 9}, {1, 9, 9}, {1, 1, 1}}, want: 5},
	{name: "先右后下更优", grid: [][]int{{1, 1, 1}, {9, 9, 1}, {9, 9, 1}}, want: 5},
	{name: "绕开大数", grid: [][]int{{1, 4, 8}, {6, 2, 7}}, want: 14},

	// 约束上限：元素最大值为 200
	{name: "2x2全200", grid: [][]int{{200, 200}, {200, 200}}, want: 600},
	{name: "3x3全200", grid: [][]int{{200, 200, 200}, {200, 200, 200}, {200, 200, 200}}, want: 1000},
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

// 验证两个版本都不修改输入 grid
func TestMinPathSum不修改输入(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	original := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	MinPathSum(grid)
	MinPathSumOptimized(grid)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != original[i][j] {
				t.Errorf("输入被修改: grid[%d][%d] = %d, want %d", i, j, grid[i][j], original[i][j])
			}
		}
	}
}

// 压力场景：200×200 最大值网格
func TestMinPathSum压力场景(t *testing.T) {
	const size = 200
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
		for j := range grid[i] {
			grid[i][j] = 200
		}
	}
	// 200×200 全为 200 时，最短路径走 200+200-1=399 个格子
	want := 399 * 200
	if got := MinPathSum(grid); got != want {
		t.Errorf("200x200全200网格: MinPathSum = %d, want %d", got, want)
	}
	if got := MinPathSumOptimized(grid); got != want {
		t.Errorf("200x200全200网格: MinPathSumOptimized = %d, want %d", got, want)
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
