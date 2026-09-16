package maximumaveragesubtree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestMaximumAverageSubtree(t *testing.T) {
	const eps = 1e-5
	tests := []struct {
		name string
		vals []int // 层序切片，math.MinInt32 表示空节点
		want float64
	}{
		// LeetCode 官方示例
		{"示例1: 左叶子最大", []int{5, 6, 1}, 6.0},
		{"示例2: 单侧链", []int{0, math.MinInt32, 1}, 1.0},

		// 边界：单节点
		{"单节点", []int{5}, 5.0},
		{"单节点为0", []int{0}, 0.0},

		// 全部节点值相同
		{"全相同值", []int{3, 3, 3, 3}, 3.0},

		// 中间子树的平均值最大：子树 2-(3,3) 平均 8/3，根子树平均 9/5
		{"中间子树最大", []int{1, 2, 0, 3, 3}, 3.0},

		// 右链 2->3->4->5，平均值逐层增大，叶子 5 最大
		{"右链", []int{2, math.MinInt32, 3, math.MinInt32, 4, math.MinInt32, 5}, 5.0},

		// 大值出现在中间节点：根 0，左子树 4-(4,4)
		{"大值在中间", []int{0, 4, 0, 4, 4}, 4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := MaximumAverageSubtree(root); math.Abs(got-tt.want) > eps {
				t.Errorf("MaximumAverageSubtree(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkMaximumAverageSubtree(b *testing.B) {
	// 构造一棵包含 10000 个节点的满二叉树（层序切片）
	vals := make([]int, 10000)
	for i := range vals {
		vals[i] = i % 1000
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaximumAverageSubtree(root)
	}
}
