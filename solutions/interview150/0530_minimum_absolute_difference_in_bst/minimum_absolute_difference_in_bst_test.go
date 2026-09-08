package minimumabsolutedifferenceinbst

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestGetMinimumDifference(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [4,2,6,1,3]", []int{4, 2, 6, 1, 3}, 1},
		{"示例2: [1,0,48,null,null,12,49]", []int{1, 0, 48, math.MinInt32, math.MinInt32, 12, 49}, 1},

		// 边界：单节点（题目保证至少 2 个节点，这里测试实现健壮性）
		{"单节点", []int{1}, math.MaxInt32},
		// 边界：两个节点
		{"两个节点", []int{2, 0}, 2},
		// 边界：完全递增的最小差在右子树
		{"差在右子树", []int{5, 3, 7, 1, 4, 6, 8}, 1},
		// 边界：最小差跨越根节点（左子树最大值与根的差），中序 3,5,7,10,15,18
		{"差跨越根", []int{10, 5, 15, 3, 7, math.MinInt32, 18}, 2},
		// 边界：最小差在左子树
		{"差在左子树", []int{100, 50, 200, 25, 75}, 25},
		// 边界：只有右子树
		{"只有右子树", []int{1, math.MinInt32, 3, math.MinInt32, 5}, 2},
		// 边界：大值差
		{"大值差", []int{100000, 0}, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := GetMinimumDifference(root); got != tt.want {
				t.Errorf("GetMinimumDifference(%v) = %d, want %d", tt.vals, got, tt.want)
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

func BenchmarkGetMinimumDifference(b *testing.B) {
	benchmarks := []struct {
		name  string
		depth int
	}{
		{"depth=100", 100},
		{"depth=1000", 1000},
		{"depth=10000", 10000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := buildIncreasingBST(bm.depth)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				GetMinimumDifference(root)
			}
		})
	}
}
