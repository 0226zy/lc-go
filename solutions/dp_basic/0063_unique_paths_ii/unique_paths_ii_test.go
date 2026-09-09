package uniquepathsii

import "testing"

var uniquePathsWithObstaclesCases = []struct {
	name string
	grid [][]int
	want int
}{
	{name: "示例1：中间有障碍", grid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, want: 2},
	{name: "示例2：右上角障碍", grid: [][]int{{0, 1}, {0, 0}}, want: 1},
	{name: "单格无障碍", grid: [][]int{{0}}, want: 1},
	{name: "单格即障碍", grid: [][]int{{1}}, want: 0},
	{name: "起点是障碍", grid: [][]int{{1, 0}}, want: 0},
	{name: "终点是障碍", grid: [][]int{{0, 0}, {0, 1}}, want: 0},
	{name: "障碍堵死通路", grid: [][]int{{0, 0}, {1, 1}, {0, 0}}, want: 0},
	{name: "单行无障碍", grid: [][]int{{0, 0, 0, 0}}, want: 1},
	{name: "单行中间障碍", grid: [][]int{{0, 0, 1, 0}}, want: 0},
	{name: "单列无障碍", grid: [][]int{{0}, {0}, {0}}, want: 1},
	{name: "障碍挡左列", grid: [][]int{{0, 0, 0}, {1, 0, 0}, {0, 0, 0}}, want: 3},
}

func TestUniquePathsWithObstacles(t *testing.T) {
	for _, tt := range uniquePathsWithObstaclesCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathsWithObstacles(tt.grid); got != tt.want {
				t.Errorf("UniquePathsWithObstacles(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestUniquePathsWithObstaclesOptimized(t *testing.T) {
	for _, tt := range uniquePathsWithObstaclesCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathsWithObstaclesOptimized(tt.grid); got != tt.want {
				t.Errorf("UniquePathsWithObstaclesOptimized(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

// 构造一个 20x20 无障碍网格用于基准测试
func benchGrid() [][]int {
	grid := make([][]int, 20)
	for i := range grid {
		grid[i] = make([]int, 20)
	}
	return grid
}

func BenchmarkUniquePathsWithObstacles(b *testing.B) {
	grid := benchGrid()
	for i := 0; i < b.N; i++ {
		UniquePathsWithObstacles(grid)
	}
}

func BenchmarkUniquePathsWithObstaclesOptimized(b *testing.B) {
	grid := benchGrid()
	for i := 0; i < b.N; i++ {
		UniquePathsWithObstaclesOptimized(grid)
	}
}
