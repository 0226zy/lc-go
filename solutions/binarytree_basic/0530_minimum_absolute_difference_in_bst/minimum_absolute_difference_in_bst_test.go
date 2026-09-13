package minimumabsolutedifferenceinbst

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestGetMinimumDifference(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序切片，math.MinInt32 表示空节点
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 平衡树 [4,2,6,1,3]", []int{4, 2, 6, 1, 3}, 1},
		{"示例2: 非平衡树 [1,0,48,null,null,12,49]", []int{1, 0, 48, math.MinInt32, math.MinInt32, 12, 49}, 1},

		// 边界：空树（题目保证至少 2 个节点，此处检验健壮性，返回哨兵 MaxInt32）
		{"空树", nil, math.MaxInt32},
		// 边界：单节点，没有可比较的前驱
		{"单节点", []int{42}, math.MaxInt32},
		// 边界：仅两个节点
		{"两个节点", []int{5, 2}, 3},
		// 边界：左斜链式树，中序为 1,2,3,4
		{"左斜链式树", []int{4, 3, math.MinInt32, 2, math.MinInt32, 1}, 1},
		// 边界：右斜链式树，中序为 1,2,3,4
		{"右斜链式树", []int{1, math.MinInt32, 2, math.MinInt32, 3, math.MinInt32, 4}, 1},
		// 边界：含负数，中序为 -10,-3,0,5
		{"含负数", []int{0, -3, 5, -10}, 3},
		// 边界：最小差出现在左子树内部，中序为 20,21,50,80,100
		{"最小差在左子树", []int{50, 21, 100, 20}, 1},
		// 边界：最小差跨越根节点，中序为 1,3,6,10,12
		{"最小差跨越根节点", []int{6, 3, 12, 1, math.MinInt32, 10}, 2},
		// 边界：相邻节点差值较大
		{"大差值", []int{100000, 0}, 100000},
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

// buildChainBST 构造一棵值为 1..n 的右斜链式 BST，用于基准测试
func buildChainBST(n int) *datastructures.TreeNode {
	var root *datastructures.TreeNode
	for i := n; i >= 1; i-- {
		root = &datastructures.TreeNode{Val: i, Right: root}
	}
	return root
}

func BenchmarkGetMinimumDifference(b *testing.B) {
	benchmarks := []struct {
		name string
		size int
	}{
		{"100节点链式树", 100},
		{"1000节点链式树", 1000},
		{"10000节点链式树", 10000},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := buildChainBST(bm.size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				GetMinimumDifference(root)
			}
		})
	}
}
