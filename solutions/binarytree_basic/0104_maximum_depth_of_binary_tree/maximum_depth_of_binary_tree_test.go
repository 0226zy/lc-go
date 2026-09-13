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
		{"示例1: 三层二叉树", []int{3, 9, 20, nilNode, nilNode, 15, 7}, 3},
		{"示例2: 只有右孩子的两节点树", []int{1, nilNode, 2}, 2},

		// 边界情况
		{"空树返回0", nil, 0},
		{"单节点树", []int{42}, 1},
		{"纯左链树", []int{1, 2, nilNode, 3, nilNode, 4}, 4},
		{"纯右链树", []int{1, nilNode, 2, nilNode, 3}, 3},
		{"含负数的树", []int{-1, -2, -3, -4}, 3},
		{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, 3},
		{"左浅右深的不平衡树", []int{1, 2, 3, nilNode, nilNode, 4, nilNode, 5}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := MaxDepth(root); got != tt.want {
				t.Errorf("MaxDepth(%v) = %d, 期望 %d", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"链式树深度10", makeChain(10)},
		{"链式树深度100", makeChain(100)},
		{"链式树深度1000", makeChain(1000)},
		{"满二叉树深度10", makeFullTree(10)},
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

// makeChain 生成长度为 n 的左链树的层序序列（每个节点只有左孩子）
func makeChain(n int) []int {
	vals := make([]int, 0, 2*n-1)
	for i := 1; i <= n; i++ {
		vals = append(vals, i)
		if i < n {
			vals = append(vals, math.MinInt32)
		}
	}
	return vals
}

// makeFullTree 生成深度为 depth 的满二叉树的层序序列（共 2^depth-1 个节点）
func makeFullTree(depth int) []int {
	n := 1<<depth - 1
	vals := make([]int, n)
	for i := range vals {
		vals[i] = i + 1
	}
	return vals
}
