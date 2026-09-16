package numberofislandsii

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestNumIslands2(t *testing.T) {
	tests := []struct {
		name      string
		m         int
		n         int
		positions [][]int
		want      []int
	}{
		// LeetCode 官方示例
		{"示例1: 3x3 逐步填海", 3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}, []int{1, 1, 2, 3}},
		{"示例2: 1x1 单格", 1, 1, [][]int{{0, 0}}, []int{1}},

		// 边界：空操作序列
		{"无操作", 2, 2, [][]int{}, []int{}},

		// 边界：重复填海不改变岛屿数
		{"重复位置", 2, 2, [][]int{{0, 0}, {0, 0}, {1, 1}, {0, 1}}, []int{1, 1, 2, 1}},

		// 典型场景：后填的陆地像桥一样把两个岛屿连成一个
		{"一行搭桥合并", 1, 3, [][]int{{0, 0}, {0, 2}, {0, 1}}, []int{1, 2, 1}},

		// 典型场景：一块陆地同时合并三个方向的岛屿
		{"中心合并三块陆地", 3, 3, [][]int{{0, 1}, {1, 0}, {1, 2}, {1, 1}}, []int{1, 2, 3, 1}},

		// 边界：互不相邻的孤立岛屿
		{"全孤立", 3, 3, [][]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}}, []int{1, 2, 3, 4}},

		// 边界：狭长网格
		{"一列依次填满", 4, 1, [][]int{{0, 0}, {3, 0}, {1, 0}, {2, 0}}, []int{1, 2, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumIslands2(tt.m, tt.n, tt.positions)
			if !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("NumIslands2(%d, %d, %v) = %v, want %v", tt.m, tt.n, tt.positions, got, tt.want)
			}
		})
	}
}

func BenchmarkNumIslands2(b *testing.B) {
	// 100x100 网格，按行优先顺序填满全部 10000 个格子
	m, n := 100, 100
	positions := make([][]int, 0, m*n)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			positions = append(positions, []int{r, c})
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NumIslands2(m, n, positions)
	}
}
