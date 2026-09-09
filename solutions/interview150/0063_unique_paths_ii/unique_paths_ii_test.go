package uniquepathsii

import "testing"

var uniquePathsWithObstaclesCases = []struct {
	name         string
	obstacleGrid [][]int
	want         int
}{
	// LeetCode 官方示例
	{name: "示例1: 3x3中间有障碍", obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, want: 2},
	{name: "示例2: 2x2右上角障碍", obstacleGrid: [][]int{{0, 1}, {0, 0}}, want: 1},

	// 边界：起点或终点是障碍
	{name: "起点是障碍返回0", obstacleGrid: [][]int{{1, 0}, {0, 0}}, want: 0},
	{name: "终点是障碍返回0", obstacleGrid: [][]int{{0, 0}, {0, 1}}, want: 0},

	// 边界：最小网格
	{name: "1x1无障碍", obstacleGrid: [][]int{{0}}, want: 1},
	{name: "1x1有障碍", obstacleGrid: [][]int{{1}}, want: 0},

	// 边界：单行单列
	{name: "1x3单行无障碍", obstacleGrid: [][]int{{0, 0, 0}}, want: 1},
	{name: "1x3单行中间障碍", obstacleGrid: [][]int{{0, 1, 0}}, want: 0},
	{name: "3x1单列无障碍", obstacleGrid: [][]int{{0}, {0}, {0}}, want: 1},
	{name: "3x1单列中间障碍", obstacleGrid: [][]int{{0}, {1}, {0}}, want: 0},

	// 边界：第一行/第一列遇到障碍后后续格子不可达
	{name: "第一行首格后全被挡", obstacleGrid: [][]int{{0, 1, 0}, {0, 0, 0}}, want: 1},
	{name: "第一列首格下全被挡", obstacleGrid: [][]int{{0, 0}, {1, 0}, {0, 0}}, want: 1},

	// 边界：障碍堵死通路
	{name: "障碍堵死通路", obstacleGrid: [][]int{{0, 0}, {1, 1}, {0, 0}}, want: 0},

	// 边界：全无障碍（退化为62题）
	{name: "2x2无障碍", obstacleGrid: [][]int{{0, 0}, {0, 0}}, want: 2},
	{name: "3x3无障碍", obstacleGrid: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, want: 6},

	// 边界：复杂障碍分布
	{name: "障碍挡左列", obstacleGrid: [][]int{{0, 0, 0}, {1, 0, 0}, {0, 0, 0}}, want: 3},
	{name: "4x4多障碍", obstacleGrid: [][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}, want: 4},
	{name: "曲折通路仅一条", obstacleGrid: [][]int{{0, 1, 1}, {0, 0, 1}, {1, 0, 0}}, want: 1},
	{name: "之字形障碍", obstacleGrid: [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}, want: 2},
}

func TestUniquePathsWithObstacles(t *testing.T) {
	for _, tt := range uniquePathsWithObstaclesCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathsWithObstacles(tt.obstacleGrid); got != tt.want {
				t.Errorf("UniquePathsWithObstacles(%v) = %d, want %d", tt.obstacleGrid, got, tt.want)
			}
		})
	}
}

func TestUniquePathsWithObstaclesOptimized(t *testing.T) {
	for _, tt := range uniquePathsWithObstaclesCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathsWithObstaclesOptimized(tt.obstacleGrid); got != tt.want {
				t.Errorf("UniquePathsWithObstaclesOptimized(%v) = %d, want %d", tt.obstacleGrid, got, tt.want)
			}
		})
	}
}

// 生成 m×n 的无障碍网格
func makeGrid(m, n int) [][]int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	return grid
}

// 压力场景：验证大网格不会 panic / 溢出
func TestUniquePathsWithObstaclesStress(t *testing.T) {
	t.Run("10x10无障碍网格", func(t *testing.T) {
		grid := makeGrid(10, 10)
		if got := UniquePathsWithObstacles(grid); got != 48620 {
			t.Errorf("UniquePathsWithObstacles(10x10无障碍) = %d, want %d", got, 48620)
		}
	})

	t.Run("100x100棋盘式障碍网格", func(t *testing.T) {
		// 斜对角放障碍，保证仍有路径且数量在 int 范围内
		grid := makeGrid(100, 100)
		for i := 1; i < 99; i++ {
			grid[i][i] = 1
		}
		if got := UniquePathsWithObstacles(grid); got <= 0 {
			t.Errorf("UniquePathsWithObstacles(100x100) = %d, 期望为正数", got)
		}
	})

	t.Run("100x100仅对角线通路", func(t *testing.T) {
		// 除第一行和最后一列外全为障碍，只剩唯一通路
		grid := makeGrid(100, 100)
		for i := 1; i < 100; i++ {
			for j := 0; j < 99; j++ {
				grid[i][j] = 1
			}
		}
		if got := UniquePathsWithObstacles(grid); got != 1 {
			t.Errorf("UniquePathsWithObstacles(100x100唯一通路) = %d, want 1", got)
		}
	})
}

func BenchmarkUniquePathsWithObstacles(b *testing.B) {
	grid := makeGrid(20, 20)
	for i := 0; i < b.N; i++ {
		UniquePathsWithObstacles(grid)
	}
}

func BenchmarkUniquePathsWithObstaclesOptimized(b *testing.B) {
	grid := makeGrid(20, 20)
	for i := 0; i < b.N; i++ {
		UniquePathsWithObstaclesOptimized(grid)
	}
}
