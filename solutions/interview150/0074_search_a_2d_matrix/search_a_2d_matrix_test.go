package searcha2dmatrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		// LeetCode 官方示例
		{"示例1: target=3", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{"示例2: target=13", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},

		// 边界：单元素矩阵
		{"单元素命中", [][]int{{5}}, 5, true},
		{"单元素未命中", [][]int{{5}}, 1, false},
		{"单元素更大目标", [][]int{{5}}, 10, false},

		// 边界：target 是矩阵最小/最大值
		{"命中最小值", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 1, true},
		{"命中最大值", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 60, true},

		// 边界：跨行边界值
		{"命中行首", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 10, true},
		{"命中行尾", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 20, true},
		{"跨行之间不存在", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 8, false},
		{"跨行之间不存在2", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 21, false},

		// 边界：单行矩阵
		{"单行命中", [][]int{{1, 3, 5}}, 3, true},
		{"单行未命中", [][]int{{1, 3, 5}}, 4, false},

		// 边界：单列矩阵
		{"单列命中", [][]int{{1}, {3}, {5}}, 3, true},
		{"单列未命中", [][]int{{1}, {3}, {5}}, 4, false},

		// 边界：负数
		{"负数矩阵", [][]int{{-10, -8, -6}, {-4, -2, 0}}, -8, true},
		{"负数矩阵未命中", [][]int{{-10, -8, -6}, {-4, -2, 0}}, -7, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("SearchMatrix(%v, %d) = %v, want %v", tt.matrix, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkSearchMatrix(b *testing.B) {
	matrix := makeMatrix(100, 100)
	benchmarks := []struct {
		name   string
		target int
	}{
		{"命中中间", 5050},
		{"命中角落", 10099},
		{"不存在", 10101},
		{"小于最小值", -1},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SearchMatrix(matrix, bm.target)
			}
		})
	}
}

// makeMatrix 生成 m 行 n 列的跨行有序矩阵
func makeMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = i*n + j
		}
	}
	return matrix
}
