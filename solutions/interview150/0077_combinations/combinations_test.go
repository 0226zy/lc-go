package combinations

import (
	"reflect"
	"testing"
)

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

		// 边界：k=n，只有一种组合
		{"k=n", 3, 3, [][]int{{1, 2, 3}}},
		// 边界：k=1，返回全部单元素组合
		{"k=1", 3, 1, [][]int{{1}, {2}, {3}}},
		// 更多验证
		{"n=5,k=3", 5, 3, [][]int{
			{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5},
			{1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Combine(tt.n, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Combine(%d, %d) = %v, want %v", tt.n, tt.k, got, tt.want)
			}
		})
	}
}

func TestCombineCount(t *testing.T) {
	// C(4,2)=6, C(20,10)=184756，用组合数验证结果数量
	if got := len(Combine(4, 2)); got != 6 {
		t.Errorf("len(Combine(4, 2)) = %d, want 6", got)
	}
	if got := len(Combine(20, 10)); got != 184756 {
		t.Errorf("len(Combine(20, 10)) = %d, want 184756", got)
	}
}

func BenchmarkCombine(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		k    int
	}{
		{"n=10,k=5", 10, 5},
		{"n=15,k=7", 15, 7},
		{"n=20,k=10", 20, 10},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Combine(bm.n, bm.k)
			}
		})
	}
}
