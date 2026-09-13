package permutationsii

import (
	"reflect"
	"sort"
	"testing"
)

// normalize 对二维切片排序，用于忽略排列结果的顺序差异
func normalize(result [][]int) [][]int {
	out := make([][]int, len(result))
	for i, p := range result {
		c := make([]int, len(p))
		copy(c, p)
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

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: nums=[1,1,2]",
			[]int{1, 1, 2},
			[][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{"示例2: nums=[1,2,3]",
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},

		// 边界：单元素
		{"单元素", []int{1}, [][]int{{1}}},
		// 边界：全部元素相同，只有一种排列
		{"全部相同", []int{2, 2, 2}, [][]int{{2, 2, 2}}},
		// 边界：含负数
		{"含负数", []int{-1, -1, 2},
			[][]int{{-1, -1, 2}, {-1, 2, -1}, {2, -1, -1}}},
		// 边界：三组重复
		{"多组重复", []int{1, 1, 2, 2},
			[][]int{{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1},
				{2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PermuteUnique(tt.nums)
			got, want := normalize(got), normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("PermuteUnique(%v) = %v, want %v", tt.nums, got, want)
			}
		})
	}
}

func BenchmarkPermuteUnique(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"长度8无重复", []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"长度8含重复", []int{1, 1, 2, 2, 3, 3, 4, 5}},
		{"长度10大量重复", []int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5}},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			nums := make([]int, len(bm.nums))
			for i := 0; i < b.N; i++ {
				copy(nums, bm.nums) // PermuteUnique 内部会排序，每次恢复原始输入
				PermuteUnique(nums)
			}
		})
	}
}
