package binarytreerightsideview

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

func TestRightSideView(t *testing.T) {
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
			name: "官方示例1",
			root: buildTree([]int{1, 2, 3, math.MinInt32, 5, math.MinInt32, 4}),
			want: []int{1, 3, 4},
		},
		{
			name: "官方示例2",
			root: buildTree([]int{1, math.MinInt32, 3}),
			want: []int{1, 3},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{1, 3, 7},
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
			name: "左子树更深",
			root: buildTree([]int{1, 2, math.MinInt32, 3, math.MinInt32, 4}),
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RightSideView(tt.root)
			if !equalSlice(got, tt.want) {
				t.Errorf("RightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRightSideViewDFS(t *testing.T) {
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
			name: "官方示例1",
			root: buildTree([]int{1, 2, 3, math.MinInt32, 5, math.MinInt32, 4}),
			want: []int{1, 3, 4},
		},
		{
			name: "官方示例2",
			root: buildTree([]int{1, math.MinInt32, 3}),
			want: []int{1, 3},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{1, 3, 7},
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
			name: "左子树更深",
			root: buildTree([]int{1, 2, math.MinInt32, 3, math.MinInt32, 4}),
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RightSideViewDFS(tt.root)
			if !equalSlice(got, tt.want) {
				t.Errorf("RightSideViewDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkRightSideView(b *testing.B) {
	root := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RightSideView(root)
	}
}

func BenchmarkRightSideViewDFS(b *testing.B) {
	root := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RightSideViewDFS(root)
	}
}
