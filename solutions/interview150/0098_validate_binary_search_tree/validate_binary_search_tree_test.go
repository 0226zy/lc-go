package validatebinarysearchtree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: 合法BST", []int{2, 1, 3}, true},
		{"示例2: 根右子树藏小值", []int{5, 1, 4, math.MinInt32, math.MinInt32, 3, 6}, false},

		// 边界：单节点
		{"单节点", []int{1}, true},
		// 边界：空树（约定为合法）
		{"空树", nil, true},
		// 边界：只有左子树且合法
		{"只有左子树合法", []int{2, 1}, true},
		// 边界：只有左子树但非法（左孩子不小于根）
		{"只有左子树非法", []int{1, 2}, false},
		// 边界：相等的值不算 BST（BST 通常要求严格小于/大于）
		{"重复值非法", []int{1, 1}, false},
		// 边界：深层违规——根的右子树中藏有小于根的值（经典陷阱）
		{"右子树藏小值", []int{3, 1, 4, math.MinInt32, math.MinInt32, 2}, false},
		// 边界：int32 极值
		{"int32最小值作根", []int{math.MinInt32}, true},
		// 边界：较大的合法 BST（13 是 14 的左孩子）
		{"较大合法BST", []int{8, 3, 10, 1, 6, math.MinInt32, 14, math.MinInt32, math.MinInt32, 4, 7, 13}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := IsValidBST(root); got != tt.want {
				t.Errorf("IsValidBST(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

// buildValidBST 构造深度为 depth 的合法平衡 BST（值域 1..2^depth-1），用于基准测试
func buildValidBST(depth int) *datastructures.TreeNode {
	var build func(lo, hi int) *datastructures.TreeNode
	build = func(lo, hi int) *datastructures.TreeNode {
		if lo > hi {
			return nil
		}
		mid := (lo + hi) / 2
		return &datastructures.TreeNode{
			Val:   mid,
			Left:  build(lo, mid-1),
			Right: build(mid+1, hi),
		}
	}
	return build(1, (1<<depth)-1)
}

func BenchmarkIsValidBST(b *testing.B) {
	benchmarks := []struct {
		name  string
		depth int
	}{
		{"depth=10(1023节点)", 10},
		{"depth=13(8191节点)", 13},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := buildValidBST(bm.depth)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				IsValidBST(root)
			}
		})
	}
}
