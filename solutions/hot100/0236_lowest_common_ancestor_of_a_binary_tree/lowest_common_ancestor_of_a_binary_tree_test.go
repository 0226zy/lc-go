package lowestcommonancestor

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// findNode 在二叉树中查找值为 val 的节点
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	// 构建示例共用树: [3,5,1,6,2,0,8,null,null,7,4]
	exampleTree := datastructures.NewTreeFromSlice([]int{
		3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4,
	})

	tests := []struct {
		name   string
		root   *TreeNode
		pVal   int
		qVal   int
		wantVal int
	}{
		{
			name:   "示例1-p和q分别在左右子树",
			root:   exampleTree,
			pVal:   5,
			qVal:   1,
			wantVal: 3,
		},
		{
			name:   "示例2-p是q的祖先",
			root:   exampleTree,
			pVal:   5,
			qVal:   4,
			wantVal: 5,
		},
		{
			name:   "示例3-只有两个节点",
			root:   datastructures.NewTreeFromSlice([]int{1, 2}),
			pVal:   1,
			qVal:   2,
			wantVal: 1,
		},
		{
			name:   "p和q都在左子树",
			root:   exampleTree,
			pVal:   6,
			qVal:   4,
			wantVal: 5,
		},
		{
			name:   "p和q都在右子树",
			root:   exampleTree,
			pVal:   0,
			qVal:   8,
			wantVal: 1,
		},
		{
			name:   "p是根节点",
			root:   exampleTree,
			pVal:   3,
			qVal:   4,
			wantVal: 3,
		},
		{
			name:   "q是根节点",
			root:   exampleTree,
			pVal:   7,
			qVal:   3,
			wantVal: 3,
		},
		{
			name:   "p和q是兄弟节点",
			root:   exampleTree,
			pVal:   6,
			qVal:   2,
			wantVal: 5,
		},
		{
			name:   "深层节点",
			root:   exampleTree,
			pVal:   7,
			qVal:   4,
			wantVal: 2,
		},
		{
			name:   "链式树",
			root:   datastructures.NewTreeFromSlice([]int{1, 2, math.MinInt32, 3, math.MinInt32, 4}),
			pVal:   4,
			qVal:   3,
			wantVal: 3,
		},
		{
			name:   "包含负数",
			root:   datastructures.NewTreeFromSlice([]int{-1, -2, -3, -4, -5}),
			pVal:   -4,
			qVal:   -5,
			wantVal: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := findNode(tt.root, tt.pVal)
			q := findNode(tt.root, tt.qVal)
			if p == nil || q == nil {
				t.Fatalf("p or q not found in tree")
			}

			got := LowestCommonAncestor(tt.root, p, q)
			if got == nil {
				t.Errorf("LowestCommonAncestor() = nil, want %d", tt.wantVal)
				return
			}
			if got.Val != tt.wantVal {
				t.Errorf("LowestCommonAncestor() = %d, want %d", got.Val, tt.wantVal)
			}
		})
	}
}

func BenchmarkLowestCommonAncestor(b *testing.B) {
	root := datastructures.NewTreeFromSlice([]int{
		3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4,
	})
	p := findNode(root, 5)
	q := findNode(root, 1)
	for i := 0; i < b.N; i++ {
		_ = LowestCommonAncestor(root, p, q)
	}
}
