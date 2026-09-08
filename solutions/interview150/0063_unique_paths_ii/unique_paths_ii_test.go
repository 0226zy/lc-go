package uniquepathsii

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name         string
		obstacleGrid [][]int
		want         int
	}{
		// LeetCode 官方示例
		{"示例1: 3x3中间有障碍", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{"示例2: 2x2右上角障碍", [][]int{{0, 1}, {0, 0}}, 1},

		// 边界：起点或终点是障碍
		{"起点是障碍返回0", [][]int{{1, 0}, {0, 0}}, 0},
		{"终点是障碍返回0", [][]int{{0, 0}, {0, 1}}, 0},

		// 边界：最小网格
		{"1x1无障碍", [][]int{{0}}, 1},
		{"1x1有障碍", [][]int{{1}}, 0},

		// 边界：单行单列
		{"1x3单行无障碍", [][]int{{0, 0, 0}}, 1},
		{"1x3单行中间障碍", [][]int{{0, 1, 0}}, 0},
		{"3x1单列无障碍", [][]int{{0}, {0}, {0}}, 1},
		{"3x1单列中间障碍", [][]int{{0}, {1}, {0}}, 0},

		// 边界：第一行/第一列遇到障碍后后续格子不可达
		{"第一行首格后全被挡", [][]int{{0, 1, 0}, {0, 0, 0}}, 1},
		{"第一列首格下全被挡", [][]int{{0, 0}, {1, 0}, {0, 0}}, 1},

		// 边界：全无障碍（退化为62题）
		{"2x2无障碍", [][]int{{0, 0}, {0, 0}}, 2},
		{"3x3无障碍", [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 6},

		// 边界：复杂障碍分布
		{"4x4多障碍", [][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}, 4},
		{"曲折通路仅一条", [][]int{{0, 1, 1}, {0, 0, 1}, {1, 0, 0}}, 1},
		{"之字形障碍", [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathsWithObstacles(tt.obstacleGrid); got != tt.want {
				t.Errorf("UniquePathsWithObstacles(%v) = %d, want %d", tt.obstacleGrid, got, tt.want)
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

// 压力场景：100x100 网格（约束上限），答案保证 <= 2*10^9 的前提是题目数据有障碍，
// 这里验证无障碍的 10x10 大网格与随机障碍的 100x100 网格不会 panic / 溢出。
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
		got := UniquePathsWithObstacles(grid)
		if got <= 0 {
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
	grid10 := makeGrid(10, 10)
	grid100 := makeGrid(100, 100)
	for i := 1; i < 99; i++ {
		grid100[i][i] = 1
	}

	benchmarks := []struct {
		name string
		grid [][]int
	}{
		{"3x3示例网格", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{"10x10无障碍", grid10},
		{"100x100棋盘式障碍", grid100},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UniquePathsWithObstacles(bm.grid)
			}
		})
	}
}
