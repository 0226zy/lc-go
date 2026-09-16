package closestbinarysearchtreevalue

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestClosestValue(t *testing.T) {
	tests := []struct {
		name   string
		vals   []int
		target float64
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: 平衡BST", []int{4, 2, 5, 1, 3}, 3.714286, 4},
		{"示例2: 单节点", []int{1}, 4.428571, 1},

		// 边界：单节点树
		{"单节点目标更小", []int{5}, 1.0, 5},
		{"单节点目标更大", []int{5}, 100.0, 5},

		// 典型场景
		{"目标恰等于某节点值", []int{4, 2, 5, 1, 3}, 2.0, 2},
		{"目标落在叶子附近", []int{4, 2, 5, 1, 3}, 1.4, 1},
		{"目标落在内部节点附近", []int{4, 2, 5, 1, 3}, 4.9, 5},
		{"目标小于所有节点值", []int{4, 2, 5, 1, 3}, -3.0, 1},
		{"目标大于所有节点值", []int{4, 2, 5, 1, 3}, 100.0, 5},

		// 边界：退化树（链表形状）
		{"右斜树", []int{1, math.MinInt32, 2, math.MinInt32, 3, math.MinInt32, 4}, 2.6, 3},
		{"左斜树", []int{4, 3, math.MinInt32, 2, math.MinInt32, 1}, 1.2, 1},

		// 边界：浮点目标恰好处于两个值中点时按路径先遇到者胜出（题目保证唯一答案，此处验证逻辑稳定性）
		{"接近根节点的目标", []int{10, 5, 15, 3, 7, 12, 20}, 6.8, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := ClosestValue(root, tt.target); got != tt.want {
				t.Errorf("ClosestValue(%v, %v) = %d, want %d", tt.vals, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkClosestValue(b *testing.B) {
	// 构造一条含 1023 个节点的右斜树，作为最长搜索路径的压力场景
	n := (1 << 10) - 1
	root := &datastructures.TreeNode{Val: 0}
	cur := root
	for i := 1; i < n; i++ {
		cur.Right = &datastructures.TreeNode{Val: i}
		cur = cur.Right
	}

	b.Run("右斜树深1023", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ClosestValue(root, float64(n)-0.5)
		}
	})
}
