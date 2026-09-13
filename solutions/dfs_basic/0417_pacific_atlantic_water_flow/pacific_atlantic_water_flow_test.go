package pacificatlantic

import (
	"reflect"
	"sort"
	"testing"
)

// normalize 对二维坐标切片排序，用于忽略结果中坐标的返回顺序
func normalize(result [][]int) [][]int {
	out := make([][]int, len(result))
	for i, cell := range result {
		c := make([]int, len(cell))
		copy(c, cell)
		sort.Ints(c)
		out[i] = c
	}
	sort.Slice(out, func(i, j int) bool {
		if len(out[i]) != len(out[j]) {
			return len(out[i]) < len(out[j])
		}
		for k := range out[i] {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return false
	})
	return out
}

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name    string
		heights [][]int
		want    [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 5x5矩阵",
			[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}},
		{"示例2: 1x1矩阵",
			[][]int{{1}},
			[][]int{{0, 0}}},

		// 边界：空矩阵
		{"空矩阵", nil, nil},
		// 边界：所有格子等高，全都能流到两个大洋
		{"等高2x2全可达",
			[][]int{{1, 1}, {1, 1}},
			[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
		// 边界：单行矩阵，每个格子都同时邻接两个大洋
		{"单行1x4",
			[][]int{{1, 2, 3, 4}},
			[][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
		// 边界：单调递增矩阵，只有右上角的峰能同时到达两边
		{"单调递增3x3",
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{{0, 2}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
		// 边界：中心最高的环形山脉，左上角两个低洼格子被高山挡住无法流入大西洋
		{"中心最高3x3",
			[][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
			[][]int{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PacificAtlantic(tt.heights)
			got, want := normalize(got), normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("PacificAtlantic(%v) = %v, want %v", tt.heights, got, want)
			}
		})
	}
}

func BenchmarkPacificAtlantic(b *testing.B) {
	// 构造 100x100 的阶梯高度矩阵作为压力测试
	const size = 100
	heights := make([][]int, size)
	for i := range heights {
		heights[i] = make([]int, size)
		for j := range heights[i] {
			heights[i][j] = i*size + j
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PacificAtlantic(heights)
	}
}
