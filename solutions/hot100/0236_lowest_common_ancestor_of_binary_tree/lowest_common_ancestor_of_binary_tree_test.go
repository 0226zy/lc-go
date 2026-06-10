package lowestcommonancestorofbinarytree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

// findNode 在树中按值查找第一个匹配的节点（假设值唯一）
func findNode(root *datastructures.TreeNode, val int) *datastructures.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := findNode(root.Left, val)
	if left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name   string
		vals   []int
		pVal   int
		qVal   int
		wantVal int
	}{
		{
			name:    "官方示例1：p=5, q=1, LCA=3",
			vals:    []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4},
			pVal:    5,
			qVal:    1,
			wantVal: 3,
		},
		{
			name:    "官方示例2：p=5, q=4, LCA=5",
			vals:    []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4},
			pVal:    5,
			qVal:    4,
			wantVal: 5,
		},
		{
			name:    "p 是 q 的祖先",
			vals:    []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4},
			pVal:    5,
			qVal:    7,
			wantVal: 5,
		},
		{
			name:    "q 是 p 的祖先",
			vals:    []int{3, 5, 1},
			pVal:    5,
			qVal:    3,
			wantVal: 3,
		},
		{
			name:    "完全二叉树：p=4, q=7, LCA=3",
			vals:    []int{3, 5, 1, 6, 2, 0, 8},
			pVal:    6,
			qVal:    0,
			wantVal: 3,
		},
		{
			name:    "单节点",
			vals:    []int{1},
			pVal:    1,
			qVal:    1,
			wantVal: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.vals)
			p := findNode(root, tt.pVal)
			q := findNode(root, tt.qVal)
			want := findNode(root, tt.wantVal)
			got := LowestCommonAncestor(root, p, q)
			if got != want {
				t.Errorf("LowestCommonAncestor() = %v, want %v", got.Val, tt.wantVal)
			}
		})
	}
}

func TestLowestCommonAncestorIterative(t *testing.T) {
	tests := []struct {
		name   string
		vals   []int
		pVal   int
		qVal   int
		wantVal int
	}{
		{
			name:    "官方示例1：p=5, q=1, LCA=3",
			vals:    []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4},
			pVal:    5,
			qVal:    1,
			wantVal: 3,
		},
		{
			name:    "官方示例2：p=5, q=4, LCA=5",
			vals:    []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4},
			pVal:    5,
			qVal:    4,
			wantVal: 5,
		},
		{
			name:    "p 是 q 的祖先",
			vals:    []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4},
			pVal:    5,
			qVal:    7,
			wantVal: 5,
		},
		{
			name:    "q 是 p 的祖先",
			vals:    []int{3, 5, 1},
			pVal:    5,
			qVal:    3,
			wantVal: 3,
		},
		{
			name:    "完全二叉树：p=4, q=7, LCA=3",
			vals:    []int{3, 5, 1, 6, 2, 0, 8},
			pVal:    6,
			qVal:    0,
			wantVal: 3,
		},
		{
			name:    "单节点",
			vals:    []int{1},
			pVal:    1,
			qVal:    1,
			wantVal: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.vals)
			p := findNode(root, tt.pVal)
			q := findNode(root, tt.qVal)
			want := findNode(root, tt.wantVal)
			got := LowestCommonAncestorIterative(root, p, q)
			if got != want {
				t.Errorf("LowestCommonAncestorIterative() = %v, want %v", got.Val, tt.wantVal)
			}
		})
	}
}

func BenchmarkLowestCommonAncestor(b *testing.B) {
	vals := []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}
	root := buildTree(vals)
	p := findNode(root, 5)
	q := findNode(root, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LowestCommonAncestor(root, p, q)
	}
}

func BenchmarkLowestCommonAncestorIterative(b *testing.B) {
	vals := []int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}
	root := buildTree(vals)
	p := findNode(root, 5)
	q := findNode(root, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LowestCommonAncestorIterative(root, p, q)
	}
}
