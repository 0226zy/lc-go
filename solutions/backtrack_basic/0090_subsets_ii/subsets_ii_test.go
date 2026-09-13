package subsetsii

import (
	"reflect"
	"sort"
	"testing"
)

// normalize 对二维切片排序，用于忽略子集结果的顺序差异
func normalize(result [][]int) [][]int {
	out := make([][]int, len(result))
	for i, combo := range result {
		c := make([]int, len(combo))
		copy(c, combo)
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

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: nums=[1,2,2]",
			[]int{1, 2, 2},
			[][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{"示例2: nums=[0]",
			[]int{0},
			[][]int{{}, {0}}},

		// 边界：无重复元素，退化为第 78 题
		{"无重复元素",
			[]int{1, 2, 3},
			[][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
		// 边界：全部元素相同
		{"全部元素相同",
			[]int{4, 4, 4},
			[][]int{{}, {4}, {4, 4}, {4, 4, 4}}},
		// 边界：包含负数且未排序
		{"含负数且乱序",
			[]int{2, 1, 2, -1},
			[][]int{{}, {-1}, {1}, {2}, {-1, 1}, {-1, 2}, {1, 2}, {2, 2}, {-1, 1, 2}, {-1, 2, 2}, {1, 2, 2}, {-1, 1, 2, 2}}},
		// 边界：两组重复元素
		{"两组重复元素",
			[]int{1, 1, 2, 2},
			[][]int{{}, {1}, {2}, {1, 1}, {1, 2}, {2, 2}, {1, 1, 2}, {1, 2, 2}, {1, 1, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubsetsWithDup(tt.nums)
			got, want := normalize(got), normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("SubsetsWithDup(%v) = %v, want %v", tt.nums, got, want)
			}
		})
	}
}

func BenchmarkSubsetsWithDup(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"n=10全重复", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{"n=10无重复", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"n=10混合重复", []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				nums := make([]int, len(bm.nums))
				copy(nums, bm.nums) // 避免排序影响基准用例本身
				SubsetsWithDup(nums)
			}
		})
	}
}
