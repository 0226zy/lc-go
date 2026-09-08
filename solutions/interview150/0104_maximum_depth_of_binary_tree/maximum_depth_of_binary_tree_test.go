package maximumdepthofbinarytree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestMaxDepth(t *testing.T) {
	nilNode := math.MinInt32
	tests := []struct {
		name string
		vals []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 三层平衡树", []int{3, 9, 20, nilNode, nilNode, 15, 7}, 3},
		{"示例2: 右斜树", []int{1, nilNode, 2}, 2},

		// 边界情况
		{"空树", nil, 0},
		{"单节点", []int{1}, 1},
		{"左斜树", []int{1, 2, nilNode, 3}, 3},
		{"两节点", []int{1, 2}, 2},
		{"负数节点", []int{-100, -200, -300}, 2},
		{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := MaxDepth(root); got != tt.want {
				t.Errorf("MaxDepth(%v) = %d, want %d", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"深度10", buildChainTree(10)},
		{"深度100", buildChainTree(100)},
		{"深度1000", buildChainTree(1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				root := datastructures.NewTreeFromSlice(bm.vals)
				MaxDepth(root)
			}
		})
	}
}

// buildChainTree 生成左斜链树的层序序列（左孩子为链，其余为 nil）
func buildChainTree(n int) []int {
	vals := make([]int, 0, 2*n-1)
	for i := 1; i <= n; i++ {
		vals = append(vals, i)
		if i < n {
			vals = append(vals, math.MinInt32)
		}
	}
	return vals
}
