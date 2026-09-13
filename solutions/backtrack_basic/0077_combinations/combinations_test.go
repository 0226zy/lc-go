package combinations

import (
	"reflect"
	"sort"
	"testing"
)

// normalize 将二维切片整体排序：每个组合内部升序，组合之间按字典序排序，
// 用于对「顺序无关」的组合结果做确定性比较。
func normalize(arr [][]int) [][]int {
	res := make([][]int, len(arr))
	for i, row := range arr {
		cp := make([]int, len(row))
		copy(cp, row)
		sort.Ints(cp)
		res[i] = cp
	}
	sort.Slice(res, func(i, j int) bool {
		a, b := res[i], res[j]
		for x := 0; x < len(a) && x < len(b); x++ {
			if a[x] != b[x] {
				return a[x] < b[x]
			}
		}
		return len(a) < len(b)
	})
	return res
}

func TestCombine(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: n=4,k=2", 4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{"示例2: n=1,k=1", 1, 1, [][]int{{1}}},

		// 边界情况
		{"边界: k等于n只有一种组合", 5, 5, [][]int{{1, 2, 3, 4, 5}}},
		{"边界: k=1返回所有单元素组合", 4, 1, [][]int{{1}, {2}, {3}, {4}}},

		// 一般情况
		{"一般: n=5,k=3", 5, 3, [][]int{
			{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5},
			{1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Combine(tt.n, tt.k)
			if !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
				t.Errorf("Combine(%d, %d) = %v, want %v", tt.n, tt.k, got, tt.want)
			}
		})
	}
}

func TestCombineResultCount(t *testing.T) {
	// 用组合数验证结果数量：C(4,2)=6，C(20,10)=184756
	cases := []struct {
		n, k, want int
	}{
		{4, 2, 6},
		{20, 10, 184756},
	}
	for _, c := range cases {
		if got := len(Combine(c.n, c.k)); got != c.want {
			t.Errorf("len(Combine(%d, %d)) = %d, want %d", c.n, c.k, got, c.want)
		}
	}
}

func BenchmarkCombine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Combine(20, 10)
	}
}
