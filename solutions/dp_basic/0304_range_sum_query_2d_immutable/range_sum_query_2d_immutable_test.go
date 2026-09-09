package rangesumquery2dimmutable

import "testing"

var leetCodeMatrix = [][]int{
	{3, 0, 1, 4, 2},
	{5, 6, 3, 2, 1},
	{1, 2, 0, 1, 5},
	{4, 1, 0, 1, 7},
	{1, 0, 3, 0, 5},
}

var sumRegionCases = []struct {
	name string
	r1   int
	c1   int
	r2   int
	c2   int
	want int
}{
	{name: "示例查询1：右下角区域", r1: 2, c1: 1, r2: 4, c2: 3, want: 8},
	{name: "示例查询2：中间区域", r1: 1, c1: 1, r2: 2, c2: 2, want: 11},
	{name: "示例查询3：右上区域", r1: 1, c1: 2, r2: 2, c2: 4, want: 12},
	{name: "整个矩阵", r1: 0, c1: 0, r2: 4, c2: 4, want: 58},
	{name: "单个元素", r1: 0, c1: 0, r2: 0, c2: 0, want: 3},
	{name: "第一行前缀", r1: 0, c1: 0, r2: 0, c2: 4, want: 10},
	{name: "第一列前缀", r1: 0, c1: 0, r2: 4, c2: 0, want: 14},
	{name: "单行中间段", r1: 2, c1: 1, r2: 2, c2: 3, want: 3},
}

func TestNumMatrix(t *testing.T) {
	nm := Constructor(leetCodeMatrix)
	for _, tt := range sumRegionCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := nm.SumRegion(tt.r1, tt.c1, tt.r2, tt.c2); got != tt.want {
				t.Errorf("SumRegion(%d, %d, %d, %d) = %d, want %d",
					tt.r1, tt.c1, tt.r2, tt.c2, got, tt.want)
			}
		})
	}
}

func TestNumMatrixSingleCell(t *testing.T) {
	// 1x1 矩阵的边界场景
	nm := Constructor([][]int{{-5}})
	if got := nm.SumRegion(0, 0, 0, 0); got != -5 {
		t.Errorf("SumRegion(0, 0, 0, 0) = %d, want %d", got, -5)
	}
}

func BenchmarkSumRegion(b *testing.B) {
	// 200x200 矩阵，约束上限
	size := 200
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
		for j := range matrix[i] {
			matrix[i][j] = i + j
		}
	}
	nm := Constructor(matrix)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nm.SumRegion(0, 0, size-1, size-1)
	}
}
