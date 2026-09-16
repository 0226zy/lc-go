package findsmallestcommonelementinallrows

import "testing"

func TestSmallestCommonElement(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: 公共元素5",
			[][]int{
				{1, 2, 3, 4, 5},
				{2, 4, 5, 8, 10},
				{3, 5, 7, 9, 11},
				{1, 3, 5, 7, 9},
			},
			5,
		},

		// 多个公共元素时取最小
		{
			"多个公共元素取最小",
			[][]int{
				{1, 2, 3},
				{2, 3, 4},
				{2, 3, 5},
			},
			2,
		},

		// 边界：无公共元素
		{
			"无公共元素返回-1",
			[][]int{
				{1, 3, 5},
				{2, 4, 6},
			},
			-1,
		},

		// 边界：单行矩阵，最小元素即该行首元素
		{
			"单行矩阵",
			[][]int{{3, 5, 7, 9}},
			3,
		},

		// 边界：单列矩阵
		{
			"单列矩阵",
			[][]int{{4}, {4}, {4}},
			4,
		},

		// 边界：所有行完全相同
		{
			"所有行完全相同",
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
			1,
		},

		// 边界：公共元素在末尾
		{
			"公共元素在末尾",
			[][]int{
				{1, 2, 3, 10000},
				{4, 5, 6, 10000},
			},
			10000,
		},

		// 典型场景：部分行包含同一较小值但非全部行
		{
			"较小值未覆盖所有行",
			[][]int{
				{1, 5},
				{2, 5},
				{3, 5},
			},
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestCommonElement(tt.mat); got != tt.want {
				t.Errorf("SmallestCommonElement(%v) = %d, want %d", tt.mat, got, tt.want)
			}
		})
	}
}

func BenchmarkSmallestCommonElement(b *testing.B) {
	// 构造 500 行 × 500 列的矩阵，行内严格递增，公共元素为最后一列
	m, n := 500, 500
	mat := make([][]int, m)
	for i := 0; i < m; i++ {
		row := make([]int, n)
		for j := 0; j < n-1; j++ {
			row[j] = (i*n+j)%9000 + 1
		}
		row[n-1] = 10000
		mat[i] = row
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SmallestCommonElement(mat)
	}
}
