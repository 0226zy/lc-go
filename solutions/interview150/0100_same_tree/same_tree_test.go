package sametree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestIsSameTree(t *testing.T) {
	nilNode := math.MinInt32
	tests := []struct {
		name string
		p    []int
		q    []int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: 完全相同", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"示例2: 结构不同", []int{1, 2}, []int{1, nilNode, 2}, false},
		{"示例3: 值不同", []int{1, 2, 1}, []int{1, 1, 2}, false},

		// 边界情况
		{"两棵空树", nil, nil, true},
		{"一空一非空", nil, []int{1}, false},
		{"单节点相同", []int{5}, []int{5}, true},
		{"单节点不同", []int{5}, []int{6}, false},
		{"深层相同树", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}, true},
		{"值相同结构不同", []int{1, 2, 3, nilNode, nilNode, 4}, []int{1, 2, 3, nilNode, 4}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := datastructures.NewTreeFromSlice(tt.p)
			q := datastructures.NewTreeFromSlice(tt.q)
			if got := IsSameTree(p, q); got != tt.want {
				t.Errorf("IsSameTree(%v, %v) = %v, want %v", tt.p, tt.q, got, tt.want)
			}
		})
	}
}

func BenchmarkIsSameTree(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"节点数7", []int{1, 2, 3, 4, 5, 6, 7}},
		{"节点数63", buildCompleteTreeVals(6)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			p := datastructures.NewTreeFromSlice(bm.vals)
			q := datastructures.NewTreeFromSlice(bm.vals)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				IsSameTree(p, q)
			}
		})
	}
}

// buildCompleteTreeVals 生成深度为 depth 的满二叉树的层序序列
func buildCompleteTreeVals(depth int) []int {
	n := 1 << depth
	vals := make([]int, n-1)
	for i := range vals {
		vals[i] = i + 1
	}
	return vals
}
