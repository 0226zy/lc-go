package lowestcommonancestorofabinarytree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// newTree 从层序切片构建二叉树，MinInt32 表示 nil
func newTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

// findNode 在树中按值查找节点（题目保证所有值互不相同）
func findNode(root *datastructures.TreeNode, val int) *datastructures.TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序表示，MinInt32 为 nil
		p, q int   // 两个目标节点的值
		want int   // 最近公共祖先的值
	}{
		// LeetCode 官方示例
		{"示例1: p=5,q=1", []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}, 5, 1, 3},
		{"示例2: p=5,q=4", []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}, 5, 4, 5},
		{"示例3: p=1,q=2", []int{1, 2}, 1, 2, 1},

		// 边界：p 是 q 的祖先
		{"p是q的祖先", []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}, 5, 7, 5},

		// 边界：q 是 p 的祖先
		{"q是p的祖先", []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}, 4, 2, 2},

		// 边界：两个节点都在右子树
		{"同在右子树", []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}, 0, 8, 1},

		// 边界：叶子节点之间
		{"两个叶子节点", []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}, 6, 4, 5},

		// 边界：最深层叶子与根
		{"深层叶子与根", []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}, 7, 3, 3},

		// 边界：负值
		{"含负值", []int{-1, -2, -3, -4, -5}, -4, -5, -2},

		// 边界：只有左子树（4→3→2→1 的左斜树）
		{"左斜树", []int{4, 3, math.MinInt32, 2, math.MinInt32, 1}, 1, 3, 3},

		// 边界：只有右子树（1→2→3 的右斜树）
		{"右斜树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, 1, 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := newTree(tt.vals)
			p := findNode(root, tt.p)
			q := findNode(root, tt.q)
			if p == nil || q == nil {
				t.Fatalf("测试数据中未找到 p=%d 或 q=%d", tt.p, tt.q)
			}
			got := LowestCommonAncestor(root, p, q)
			if got == nil || got.Val != tt.want {
				t.Errorf("LowestCommonAncestor(p=%d, q=%d) = %v, want %d", tt.p, tt.q, nodeVal(got), tt.want)
			}
		})
	}
}

// nodeVal 返回节点值，nil 节点返回 MinInt32 便于打印
func nodeVal(node *datastructures.TreeNode) int {
	if node == nil {
		return math.MinInt32
	}
	return node.Val
}

func BenchmarkLowestCommonAncestor(b *testing.B) {
	// 构建一棵较大的树：根 + 左右两个大子树，p 在左、q 在右
	root := newTree([]int{100,
		50, 150,
		25, 75, 125, 175,
		12, 37, 62, 87, 112, 137, 162, 187,
		6, 18, 31, 43, 56, 68, 81, 93, 106, 118, 131, 143, 156, 168, 181, 193,
		math.MinInt32, math.MinInt32, 8, math.MinInt32, math.MinInt32, math.MinInt32, 40})
	p := findNode(root, 8)
	q := findNode(root, 40)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LowestCommonAncestor(root, p, q)
	}
}
