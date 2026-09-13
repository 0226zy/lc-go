package combinationsumiii

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

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name string
		k    int
		n    int
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: k=3, n=7", 3, 7, [][]int{{1, 2, 4}}},
		{"示例2: k=3, n=9", 3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{"示例3: k=4, n=1", 4, 1, nil},

		// 边界：只选 1 个数，且恰好存在
		{"k=1且存在", 1, 5, [][]int{{5}}},
		// 边界：只选 1 个数，但 n 超出 1-9 范围
		{"k=1且n超出范围", 1, 10, nil},
		// 边界：k 个最小数之和已超过 n（k=4 时最小和为 1+2+3+4=10）
		{"k=4且n=9无解", 4, 9, nil},
		// 边界：选满 9 个数，和固定为 45
		{"k=9且n=45", 9, 45, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}},
		// 边界：选满 9 个数但和不为 45，无解
		{"k=9且n=44无解", 9, 44, nil},
		// 常规：多解情况
		{"k=2且n=10", 2, 10, [][]int{{1, 9}, {2, 8}, {3, 7}, {4, 6}}},
		// 常规：较大的 k
		{"k=4且n=24", 4, 24,
			[][]int{{1, 6, 8, 9}, {2, 5, 8, 9}, {2, 6, 7, 9}, {3, 4, 8, 9},
				{3, 5, 7, 9}, {3, 6, 7, 8}, {4, 5, 6, 9}, {4, 5, 7, 8}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombinationSum3(tt.k, tt.n)
			got, want := normalize(got), normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("CombinationSum3(%d, %d) = %v, want %v", tt.k, tt.n, got, want)
			}
		})
	}
}

func BenchmarkCombinationSum3(b *testing.B) {
	benchmarks := []struct {
		name string
		k    int
		n    int
	}{
		{"k=3且n=9", 3, 9},
		{"k=5且n=30", 5, 30},
		{"k=7且n=35", 7, 35},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CombinationSum3(bm.k, bm.n)
			}
		})
	}
}
