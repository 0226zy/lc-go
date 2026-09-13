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
		{"示例1: 两棵树完全相同", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"示例2: 左右子节点位置错位", []int{1, 2}, []int{1, nilNode, 2}, false},
		{"示例3: 结构相同但节点值不同", []int{1, 2, 1}, []int{1, 1, 2}, false},

		// 边界情况
		{"两棵空树", nil, nil, true},
		{"一棵空树一棵非空", nil, []int{1}, false},
		{"单节点值相同", []int{0}, []int{0}, true},
		{"单节点值不同", []int{-1}, []int{1}, false},
		{"含负数值相同", []int{-3, -1, -4}, []int{-3, -1, -4}, true},
		{"含负数值不同", []int{-3, -1}, []int{-3, -2}, false},
		{"链式树左链相同", []int{1, 2, nilNode, 3}, []int{1, 2, nilNode, 3}, true},
		{"链式树一左一右", []int{1, 2, nilNode, 3}, []int{1, nilNode, 2, nilNode, 3}, false},
		{"满二叉树完全相同", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}, true},
		{"满二叉树叶子缺失一个", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6}, false},
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
		{"深度3的满二叉树", buildCompleteTreeVals(3)},
		{"深度10的满二叉树", buildCompleteTreeVals(10)},
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
