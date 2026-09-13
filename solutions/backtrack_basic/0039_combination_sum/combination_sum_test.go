package combinationsum

import (
	"reflect"
	"sort"
	"testing"
)

// normalize 将每个组合内部排序、再对组合列表排序，便于做无序比较
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
		{"示例1: [2,3,6,7] 凑 7",
			[]int{2, 3, 6, 7}, 7,
			[][]int{{2, 2, 3}, {7}}},
		{"示例2: [2,3,5] 凑 8",
			[]int{2, 3, 5}, 8,
			[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"示例3: [2] 凑 1 无解",
			[]int{2}, 1,
			nil},

		// 边界情况
		{"单个元素恰好等于target", []int{2}, 2, [][]int{{2}}},
		{"单个元素重复多次", []int{3}, 9, [][]int{{3, 3, 3}}},
		{"包含1时组合数量最多", []int{1, 2}, 3, [][]int{{1, 1, 1}, {1, 2}}},
		{"全部元素都大于target", []int{5, 8}, 4, nil},
		{"输入未排序也能正确处理", []int{8, 7, 4, 3}, 7, [][]int{{7}, {3, 4}}},
		{"target为1且最小元素为2", []int{2, 3}, 1, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombinationSum(tt.candidates, tt.target)
			if !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
				t.Errorf("CombinationSum(%v, %d) = %v, 期望 %v",
					tt.candidates, tt.target, got, tt.want)
			}
		})
	}
}

// TestCombinationSum不修改输入 验证函数不会改动调用方传入的切片
func TestCombinationSum不修改输入(t *testing.T) {
	candidates := []int{7, 3, 2, 6}
	origin := make([]int, len(candidates))
	copy(origin, candidates)

	CombinationSum(candidates, 7)

	if !reflect.DeepEqual(candidates, origin) {
		t.Errorf("输入切片被修改: 调用前 %v, 调用后 %v", origin, candidates)
	}
}

func BenchmarkCombinationSum(b *testing.B) {
	candidates := []int{2, 3, 5, 7, 11}
	target := 30
	for i := 0; i < b.N; i++ {
		CombinationSum(candidates, target)
	}
}
