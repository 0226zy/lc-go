package kthsmallestelementinabst

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		k    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: k=1", []int{3, 1, 4, math.MinInt32, 2}, 1, 1},
		{"示例2: k=3", []int{5, 3, 6, 2, 4, math.MinInt32, math.MinInt32, 1}, 3, 3},

		// 边界：k 为节点总数（最后一个）
		{"k为总数", []int{3, 1, 4}, 3, 4},
		// 边界：k=1 且为单节点
		{"单节点k=1", []int{1}, 1, 1},
		// 边界：k 落在右子树
		{"k在右子树", []int{3, 1, 4, math.MinInt32, 2}, 4, 4},
		// 边界：左斜树（退化链）
		{"左斜树k=2", []int{4, 3, math.MinInt32, 2, math.MinInt32, 1}, 2, 2},
		// 边界：右斜树
		{"右斜树k=3", []int{1, math.MinInt32, 2, math.MinInt32, 3}, 3, 3},
		// 边界：大 k
		{"k=7满树", []int{4, 2, 6, 1, 3, 5, 7}, 7, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := KthSmallest(root, tt.k); got != tt.want {
				t.Errorf("KthSmallest(%v, %d) = %d, want %d", tt.vals, tt.k, got, tt.want)
			}
		})
	}
}

// buildIncreasingBST 构造深度为 depth 的右斜 BST（1,2,3,...,depth），用于基准测试
func buildIncreasingBST(depth int) *datastructures.TreeNode {
	var root *datastructures.TreeNode
	for i := 1; i <= depth; i++ {
		root = &datastructures.TreeNode{Val: i, Right: root}
	}
	return root
}

func BenchmarkKthSmallest(b *testing.B) {
	benchmarks := []struct {
		name  string
		depth int
		k     int
	}{
		{"depth=1000,k=500", 1000, 500},
		{"depth=10000,k=5000", 10000, 5000},
		{"depth=10000,k=10000", 10000, 10000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := buildIncreasingBST(bm.depth)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				KthSmallest(root, bm.k)
			}
		})
	}
}
