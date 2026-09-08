package minimumpathsum

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 3x3网格", [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{"示例2: 2x3网格", [][]int{{1, 2, 3}, {4, 5, 6}}, 12},

		// 边界：单元素与单行/单列
		{"单元素网格", [][]int{{5}}, 5},
		{"单行网格只能向右", [][]int{{1, 2, 3, 4}}, 10},
		{"单列网格只能向下", [][]int{{1}, {2}, {3}, {4}}, 10},

		// 边界：全零网格
		{"全零2x2网格", [][]int{{0, 0}, {0, 0}}, 0},
		{"起点为0", [][]int{{0, 0}, {1, 1}}, 1},

		// 普通场景：绕开高代价区域
		{"绕开中间大数", [][]int{{1, 100, 1}, {1, 100, 1}, {1, 1, 1}}, 5},
		{"先下后右更优", [][]int{{1, 9, 9}, {1, 9, 9}, {1, 1, 1}}, 5},
		{"先右后下更优", [][]int{{1, 1, 1}, {9, 9, 1}, {9, 9, 1}}, 5},

		// 约束上限：元素最大值为 200
		{"2x2全200", [][]int{{200, 200}, {200, 200}}, 600},
		{"3x3全200", [][]int{{200, 200, 200}, {200, 200, 200}, {200, 200, 200}}, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPathSum(tt.grid); got != tt.want {
				t.Errorf("MinPathSum(%v) = %v, want %v", tt.grid, got, tt.want)
			}
		})
	}
}

// 验证不修改输入 grid
func TestMinPathSum不修改输入(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	original := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	MinPathSum(grid)

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
}

func BenchmarkMinPathSum(b *testing.B) {
	// 小规模网格
	small := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	// 大规模网格：200×200，值为行列下标之和取模
	const size = 200
	large := make([][]int, size)
	for i := range large {
		large[i] = make([]int, size)
		for j := range large[i] {
			large[i][j] = (i + j) % 200
		}
	}

	b.Run("3x3网格", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MinPathSum(small)
		}
	})
	b.Run("200x200网格", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MinPathSum(large)
		}
	})
}
