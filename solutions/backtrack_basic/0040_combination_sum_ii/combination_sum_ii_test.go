package combinationsumii

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

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		// LeetCode 官方示例
		{"示例1: candidates=[10,1,2,7,6,1,5],target=8",
			[]int{10, 1, 2, 7, 6, 1, 5}, 8,
			[][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{"示例2: candidates=[2,5,2,1,2],target=5",
			[]int{2, 5, 2, 1, 2}, 5,
			[][]int{{1, 2, 2}, {5}}},

		// 边界：无解
		{"无解: candidates=[2],target=1", []int{2}, 1, nil},
		// 边界：单个元素恰好等于 target
		{"单元素恰好等于target", []int{5}, 5, [][]int{{5}}},
		// 边界：所有元素都是重复值，每个只能用一次
		{"全为重复值", []int{1, 1, 1, 1}, 2, [][]int{{1, 1}}},
		// 边界：大数提前剪枝
		{"大数剪枝", []int{8, 7, 4, 3}, 7, [][]int{{7}, {3, 4}}},
		// 边界：重复值跨越同层去重
		{"重复值去重", []int{2, 2, 2}, 4, [][]int{{2, 2}}},
		// 边界：每个元素只能用一次，不能重复选
		{"不能重复选取", []int{1, 2}, 3, [][]int{{1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombinationSum2(tt.candidates, tt.target)
			got, want := normalize(got), normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("CombinationSum2(%v, %d) = %v, want %v", tt.candidates, tt.target, got, want)
			}
		})
	}
}

func BenchmarkCombinationSum2(b *testing.B) {
	benchmarks := []struct {
		name       string
		candidates []int
		target     int
	}{
		{"target=15", []int{2, 3, 5, 7, 2, 3}, 15},
		{"target=25", []int{2, 3, 5, 7, 11, 2, 3, 5}, 25},
		{"target=40", []int{2, 3, 5, 7, 11, 2, 3, 5, 7}, 40},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CombinationSum2(bm.candidates, bm.target)
			}
		})
	}
}
