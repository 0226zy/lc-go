package minimumfallingpathsum

import "testing"

var minFallingPathSumCases = []struct {
	name   string
	matrix [][]int
	want   int
}{
	{name: "示例1", matrix: [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, want: 13},
	{name: "示例2含负数", matrix: [][]int{{-19, 57}, {-40, -5}}, want: -59},
	{name: "单元素", matrix: [][]int{{5}}, want: 5},
	{name: "单个负数", matrix: [][]int{{-100}}, want: -100},
	{name: "直线下降最优", matrix: [][]int{{1, 9, 9}, {9, 1, 9}, {9, 9, 1}}, want: 3},
	{name: "必须拐弯", matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, want: 12},
	{name: "全零", matrix: [][]int{{0, 0}, {0, 0}}, want: 0},
	{name: "左边缘走法", matrix: [][]int{{-80, -13, 22}, {83, 94, -5}, {73, -48, 61}}, want: -66},
}

func TestMinFallingPathSum(t *testing.T) {
	for _, tt := range minFallingPathSumCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinFallingPathSum(tt.matrix); got != tt.want {
				t.Errorf("MinFallingPathSum(%v) = %d, want %d", tt.matrix, got, tt.want)
			}
		})
	}
}

func TestMinFallingPathSumOptimized(t *testing.T) {
	for _, tt := range minFallingPathSumCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinFallingPathSumOptimized(tt.matrix); got != tt.want {
				t.Errorf("MinFallingPathSumOptimized(%v) = %d, want %d", tt.matrix, got, tt.want)
			}
		})
	}
}

func BenchmarkMinFallingPathSum(b *testing.B) {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		MinFallingPathSum(matrix)
	}
}

func BenchmarkMinFallingPathSumOptimized(b *testing.B) {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		MinFallingPathSumOptimized(matrix)
	}
}
