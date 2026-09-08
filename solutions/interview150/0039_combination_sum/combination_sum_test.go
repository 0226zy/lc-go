package combinationsum

import (
	"reflect"
	"sort"
	"testing"
)

// normalize 对二维切片排序，用于忽略组合结果的顺序差异
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

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		// LeetCode 官方示例
		{"示例1: candidates=[2,3,6,7],target=7",
			[]int{2, 3, 6, 7}, 7,
			[][]int{{2, 2, 3}, {7}}},
		{"示例2: candidates=[2,3,5],target=8",
			[]int{2, 3, 5}, 8,
			[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"示例3: candidates=[2],target=1",
			[]int{2}, 1,
			nil},

		// 边界：恰好等于 target
		{"恰好等于target", []int{2}, 2, [][]int{{2}}},
		// 边界：全是 1
		{"全为1", []int{1, 2}, 3, [][]int{{1, 1, 1}, {1, 2}}},
		// 边界：大数提前剪枝
		{"大数剪枝", []int{8, 7, 4, 3}, 7, [][]int{{7}, {3, 4}}},
		// 更多验证
		{"target=4,candidates=[1,2]", []int{1, 2}, 4,
			[][]int{{1, 1, 1, 1}, {1, 1, 2}, {2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombinationSum(tt.candidates, tt.target)
			got, want := normalize(got), normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("CombinationSum(%v, %d) = %v, want %v", tt.candidates, tt.target, got, want)
			}
		})
	}
}

func BenchmarkCombinationSum(b *testing.B) {
	benchmarks := []struct {
		name       string
		candidates []int
		target     int
	}{
		{"target=15", []int{2, 3, 5, 7}, 15},
		{"target=25", []int{2, 3, 5, 7, 11}, 25},
		{"target=40", []int{2, 3, 5, 7, 11}, 40},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CombinationSum(bm.candidates, bm.target)
			}
		})
	}
}
