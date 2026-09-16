package missingranges

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestFindMissingRanges(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		lower int
		upper int
		want  [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 中间与两端均有缺失", []int{0, 1, 3, 50, 75}, 0, 99, [][]int{{2, 2}, {4, 49}, {51, 74}, {76, 99}}},
		{"示例2: 恰好覆盖无缺失", []int{-1}, -1, -1, [][]int{}},

		// 边界：空数组
		{"空数组整段缺失", []int{}, 1, 1, [][]int{{1, 1}}},
		{"空数组长区间", []int{}, 0, 99, [][]int{{0, 99}}},

		// 边界：单元素
		{"单元素在区间正中", []int{5}, 0, 10, [][]int{{0, 4}, {6, 10}}},
		{"单元素贴左端", []int{0}, 0, 5, [][]int{{1, 5}}},
		{"单元素贴右端", []int{5}, 0, 5, [][]int{{0, 4}}},

		// 边界：负数范围
		{"负数区间", []int{-5, -2}, -7, 0, [][]int{{-7, -6}, {-4, -3}, {-1, 0}}},

		// 典型场景
		{"连续缺失一格", []int{0, 2, 4}, 0, 4, [][]int{{1, 1}, {3, 3}}},
		{"完全覆盖", []int{0, 1, 2}, 0, 2, [][]int{}},
		{"仅首段缺失", []int{5, 6, 7}, 0, 7, [][]int{{0, 4}}},
		{"仅尾段缺失", []int{0, 1, 2}, 0, 5, [][]int{{3, 5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMissingRanges(tt.nums, tt.lower, tt.upper)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("FindMissingRanges(%v, %d, %d) = %v, want %v", tt.nums, tt.lower, tt.upper, got, tt.want)
			}
		})
	}
}

func BenchmarkFindMissingRanges(b *testing.B) {
	// 构造 100 个间隔为 2 的元素，产生 100+ 个缺失区间
	nums := make([]int, 0, 100)
	for i := 0; i < 200; i += 2 {
		nums = append(nums, i)
	}

	b.Run("100个元素间隔缺失", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMissingRanges(nums, 0, 200)
		}
	})
}
