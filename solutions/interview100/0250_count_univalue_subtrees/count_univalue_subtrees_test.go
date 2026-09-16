package countunivaluesubtrees

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestCountUnivalSubtrees(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序切片，math.MinInt32 表示空位
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: 混合值",
			[]int{5, 1, 5, 5, 5, math.MinInt32, 5},
			4,
		},
		{
			"示例2: 空树",
			[]int{},
			0,
		},
		{
			"示例3: 全同值",
			[]int{5, 5, 5, 5, 5, math.MinInt32, 5},
			6,
		},

		// 边界：单节点必然是一棵同值子树
		{
			"单节点",
			[]int{1},
			1,
		},

		// 边界：完全二叉树且值全同
		{
			"满树全同值",
			[]int{1, 1, 1, 1, 1, 1, 1},
			7,
		},

		// 典型：同值但子树不同值（父同值、子异值）
		{
			"仅叶子同值",
			[]int{1, 2, 3},
			2,
		},

		// 典型：负值
		{
			"负值同值子树",
			[]int{-5, -5, -5},
			3,
		},

		// 典型：链式退化的同值树
		{
			"链式同值",
			[]int{2, 2, math.MinInt32, 2},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := CountUnivalSubtrees(root); got != tt.want {
				t.Errorf("CountUnivalSubtrees(%v) = %d, want %d", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkCountUnivalSubtrees(b *testing.B) {
	// 构造一棵 1000 节点的同值满树
	vals := make([]int, 1000)
	for i := range vals {
		vals[i] = 5
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountUnivalSubtrees(root)
	}
}
