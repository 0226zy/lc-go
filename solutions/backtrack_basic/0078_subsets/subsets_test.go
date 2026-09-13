package subsets

import (
	"reflect"
	"sort"
	"testing"
)

// normalize 对二维切片排序，用于忽略子集结果中子集顺序的差异
func normalize(result [][]int) [][]int {
	out := make([][]int, len(result))
	for i, subset := range result {
		s := make([]int, len(subset))
		copy(s, subset)
		sort.Ints(s)
		out[i] = s
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

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		// LeetCode 官方示例
		{"官方示例1: nums=[1,2,3]",
			[]int{1, 2, 3},
			[][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{"官方示例2: nums=[0]",
			[]int{0},
			[][]int{{}, {0}}},

		// 边界：单元素（含负数）
		{"单元素负数", []int{-1}, [][]int{{}, {-1}}},
		// 边界：两个元素
		{"两个元素", []int{1, 2}, [][]int{{}, {1}, {2}, {1, 2}}},
		// 边界：包含负数与零
		{"含负数和零", []int{-10, 0, 10},
			[][]int{{}, {-10}, {0}, {10}, {-10, 0}, {-10, 10}, {0, 10}, {-10, 0, 10}}},
		// 边界：输入顺序乱序（元素仍互不相同）
		{"输入乱序", []int{3, 1, 2},
			[][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subsets(tt.nums)
			got, want := normalize(got), normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Subsets(%v) = %v, want %v", tt.nums, got, want)
			}
		})
	}
}

func BenchmarkSubsets(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		Subsets(nums)
	}
}
