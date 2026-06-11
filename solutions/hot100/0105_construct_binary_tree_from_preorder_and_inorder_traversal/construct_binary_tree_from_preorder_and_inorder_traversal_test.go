package buildtreeprein

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// treeEqual 递归比较两棵二叉树是否完全相同
func treeEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && treeEqual(a.Left, b.Left) && treeEqual(a.Right, b.Right)
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			name:     "示例1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want: datastructures.NewTreeFromSlice([]int{
				3, 9, 20, math.MinInt32, math.MinInt32, 15, 7,
			}),
		},
		{
			name:     "示例2-单节点",
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     datastructures.NewTreeFromSlice([]int{-1}),
		},
		{
			name:     "只有左子树",
			preorder: []int{1, 2, 3},
			inorder:  []int{3, 2, 1},
			want: datastructures.NewTreeFromSlice([]int{
				1, 2, math.MinInt32, 3,
			}),
		},
		{
			name:     "只有右子树",
			preorder: []int{1, 2, 3},
			inorder:  []int{1, 2, 3},
			want: datastructures.NewTreeFromSlice([]int{
				1, math.MinInt32, 2, math.MinInt32, 3,
			}),
		},
		{
			name:     "完全二叉树",
			preorder: []int{1, 2, 4, 5, 3, 6, 7},
			inorder:  []int{4, 2, 5, 1, 6, 3, 7},
			want: datastructures.NewTreeFromSlice([]int{
				1, 2, 3, 4, 5, 6, 7,
			}),
		},
		{
			name:     "左右子树不平衡",
			preorder: []int{3, 9, 10, 20, 15, 7},
			inorder:  []int{10, 9, 3, 15, 20, 7},
			want: datastructures.NewTreeFromSlice([]int{
				3, 9, 20, 10, math.MinInt32, 15, 7,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildTree(tt.preorder, tt.inorder)
			if !treeEqual(got, tt.want) {
				t.Errorf("BuildTree() produced different tree")
			}
		})
	}
}

func BenchmarkBuildTree(b *testing.B) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	for i := 0; i < b.N; i++ {
		_ = BuildTree(preorder, inorder)
	}
}
