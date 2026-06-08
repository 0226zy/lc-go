package binarytree

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

func TestTreeTraversal_PreorderRecursive(t *testing.T) {
	tt := NewTreeTraversal()
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want []int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: []int{1},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "只有左子树非完全",
			root: buildTree([]int{1, 2, math.MinInt32, math.MinInt32, 5}),
			want: []int{1, 2, 5},
		},
		{
			name: "只有右子树非完全",
			root: buildTree([]int{1, math.MinInt32, 3, math.MinInt32, 6}),
			want: []int{1, 3, 6},
		},
		{
			name: "不对称树",
			root: buildTree([]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 7}),
			want: []int{1, 2, 4, 3, 7},
		},
		{
			name: "四层完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			want: []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tt.PreorderRecursive(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PreorderRecursive() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTreeTraversal_PreorderIterative(t *testing.T) {
	tt := NewTreeTraversal()
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want []int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: []int{1},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "只有左子树非完全",
			root: buildTree([]int{1, 2, math.MinInt32, math.MinInt32, 5}),
			want: []int{1, 2, 5},
		},
		{
			name: "只有右子树非完全",
			root: buildTree([]int{1, math.MinInt32, 3, math.MinInt32, 6}),
			want: []int{1, 3, 6},
		},
		{
			name: "不对称树",
			root: buildTree([]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 7}),
			want: []int{1, 2, 4, 3, 7},
		},
		{
			name: "四层完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			want: []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tt.PreorderIterative(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PreorderIterative() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTreeTraversal_InorderRecursive(t *testing.T) {
	tt := NewTreeTraversal()
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want []int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: []int{1},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: []int{3, 2, 1},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "只有左子树非完全",
			root: buildTree([]int{1, 2, math.MinInt32, math.MinInt32, 5}),
			want: []int{2, 5, 1},
		},
		{
			name: "只有右子树非完全",
			root: buildTree([]int{1, math.MinInt32, 3, math.MinInt32, 6}),
			want: []int{1, 3, 6},
		},
		{
			name: "不对称树",
			root: buildTree([]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 7}),
			want: []int{4, 2, 1, 3, 7},
		},
		{
			name: "四层完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			want: []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tt.InorderRecursive(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("InorderRecursive() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTreeTraversal_InorderIterative(t *testing.T) {
	tt := NewTreeTraversal()
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want []int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: []int{1},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: []int{3, 2, 1},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "只有左子树非完全",
			root: buildTree([]int{1, 2, math.MinInt32, math.MinInt32, 5}),
			want: []int{2, 5, 1},
		},
		{
			name: "只有右子树非完全",
			root: buildTree([]int{1, math.MinInt32, 3, math.MinInt32, 6}),
			want: []int{1, 3, 6},
		},
		{
			name: "不对称树",
			root: buildTree([]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 7}),
			want: []int{4, 2, 1, 3, 7},
		},
		{
			name: "四层完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			want: []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tt.InorderIterative(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("InorderIterative() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTreeTraversal_PostorderRecursive(t *testing.T) {
	tt := NewTreeTraversal()
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want []int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: []int{1},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: []int{3, 2, 1},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: []int{3, 2, 1},
		},
		{
			name: "只有左子树非完全",
			root: buildTree([]int{1, 2, math.MinInt32, math.MinInt32, 5}),
			want: []int{5, 2, 1},
		},
		{
			name: "只有右子树非完全",
			root: buildTree([]int{1, math.MinInt32, 3, math.MinInt32, 6}),
			want: []int{6, 3, 1},
		},
		{
			name: "不对称树",
			root: buildTree([]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 7}),
			want: []int{4, 2, 7, 3, 1},
		},
		{
			name: "四层完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			want: []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tt.PostorderRecursive(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PostorderRecursive() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTreeTraversal_PostorderIterative(t *testing.T) {
	tt := NewTreeTraversal()
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want []int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: []int{1},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: []int{3, 2, 1},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: []int{3, 2, 1},
		},
		{
			name: "只有左子树非完全",
			root: buildTree([]int{1, 2, math.MinInt32, math.MinInt32, 5}),
			want: []int{5, 2, 1},
		},
		{
			name: "只有右子树非完全",
			root: buildTree([]int{1, math.MinInt32, 3, math.MinInt32, 6}),
			want: []int{6, 3, 1},
		},
		{
			name: "不对称树",
			root: buildTree([]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 7}),
			want: []int{4, 2, 7, 3, 1},
		},
		{
			name: "四层完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			want: []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tt.PostorderIterative(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PostorderIterative() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTreeTraversal_LevelOrder(t *testing.T) {
	tt := NewTreeTraversal()
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want []int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: []int{1},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "只有左子树非完全",
			root: buildTree([]int{1, 2, math.MinInt32, math.MinInt32, 5}),
			want: []int{1, 2, 5},
		},
		{
			name: "只有右子树非完全",
			root: buildTree([]int{1, math.MinInt32, 3, math.MinInt32, 6}),
			want: []int{1, 3, 6},
		},
		{
			name: "不对称树",
			root: buildTree([]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 7}),
			want: []int{1, 2, 3, 4, 7},
		},
		{
			name: "四层完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tt.LevelOrder(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTreeTraversal_LevelOrderByLevel(t *testing.T) {
	tt := NewTreeTraversal()
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want [][]int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: [][]int{{1}},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "只有左子树非完全",
			root: buildTree([]int{1, 2, math.MinInt32, math.MinInt32, 5}),
			want: [][]int{{1}, {2}, {5}},
		},
		{
			name: "只有右子树非完全",
			root: buildTree([]int{1, math.MinInt32, 3, math.MinInt32, 6}),
			want: [][]int{{1}, {3}, {6}},
		},
		{
			name: "不对称树",
			root: buildTree([]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 7}),
			want: [][]int{{1}, {2, 3}, {4, 7}},
		},
		{
			name: "四层完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tt.LevelOrderByLevel(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("LevelOrderByLevel() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTreeTraversal_LevelOrderZigzag(t *testing.T) {
	tt := NewTreeTraversal()
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want [][]int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: [][]int{{1}},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "只有左子树非完全",
			root: buildTree([]int{1, 2, math.MinInt32, math.MinInt32, 5}),
			want: [][]int{{1}, {2}, {5}},
		},
		{
			name: "只有右子树非完全",
			root: buildTree([]int{1, math.MinInt32, 3, math.MinInt32, 6}),
			want: [][]int{{1}, {3}, {6}},
		},
		{
			name: "不对称树",
			root: buildTree([]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 7}),
			want: [][]int{{1}, {3, 2}, {4, 7}},
		},
		{
			name: "四层完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
			want: [][]int{{1}, {3, 2}, {4, 5, 6, 7}, {15, 14, 13, 12, 11, 10, 9, 8}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tt.LevelOrderZigzag(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("LevelOrderZigzag() = %v, want %v", got, tc.want)
			}
		})
	}
}
