package kthsmallest

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序切片，math.MinInt32 表示空节点
		k    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 平衡树取最小值", []int{3, 1, 4, math.MinInt32, 2}, 1, 1},
		{"示例2: 多层树取第3小", []int{5, 3, 6, 2, 4, math.MinInt32, math.MinInt32, 1}, 3, 3},

		// 边界：单节点树
		{"单节点树k=1", []int{1}, 1, 1},
		// 边界：k 等于节点总数，答案是最大值
		{"k等于节点总数", []int{3, 1, 4, math.MinInt32, 2}, 4, 4},
		// 边界：左斜链式树（中序从最底端开始）
		{"左斜链k=1", []int{3, 2, math.MinInt32, 1}, 1, 1},
		{"左斜链k=3", []int{3, 2, math.MinInt32, 1}, 3, 3},
		// 边界：右斜链式树（中序就是层序顺序）
		{"右斜链k=2", []int{1, math.MinInt32, 2, math.MinInt32, 3}, 2, 2},
		// 边界：含负数的 BST
		{"含负数k=2", []int{0, -3, 5, -5, -1, 2, 8}, 2, -3},
		// 边界：完全二叉搜索树取中间值
		{"满树取中位数", []int{4, 2, 6, 1, 3, 5, 7}, 4, 4},
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

// 空树单独测试：k 无效时不应 panic，返回零值
func TestKthSmallest空树(t *testing.T) {
	if got := KthSmallest(nil, 1); got != 0 {
		t.Errorf("KthSmallest(nil, 1) = %d, want 0", got)
	}
}

func BenchmarkKthSmallest(b *testing.B) {
	// 构造 1->2->...->n 的右斜链（最坏情况，树高等于节点数），压测递归深度与提前终止
	const n = 10000
	vals := make([]int, 0, 2*n)
	for i := 1; i <= n; i++ {
		vals = append(vals, i, math.MinInt32) // 每个节点左孩子为空、右孩子为下一个数
	}
	root := datastructures.NewTreeFromSlice(vals)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KthSmallest(root, n/2)
	}
}
