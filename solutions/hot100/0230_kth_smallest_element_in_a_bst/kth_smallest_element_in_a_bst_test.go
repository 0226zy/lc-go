package kthsmallestelementinabst

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		k    int
		want int
	}{
		{
			name: "官方示例1",
			root: buildTree([]int{3, 1, 4, math.MinInt32, 2}),
			k:    1,
			want: 1,
		},
		{
			name: "官方示例2",
			root: buildTree([]int{5, 3, 6, 2, 4, math.MinInt32, math.MinInt32, 1}),
			k:    3,
			want: 3,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			k:    1,
			want: 1,
		},
		{
			name: "三个节点-找最大",
			root: buildTree([]int{2, 1, 3}),
			k:    3,
			want: 3,
		},
		{
			name: "三个节点-找中间",
			root: buildTree([]int{2, 1, 3}),
			k:    2,
			want: 2,
		},
		{
			name: "左斜树",
			root: buildTree([]int{3, 2, math.MinInt32, 1}),
			k:    2,
			want: 2,
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			k:    3,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KthSmallest(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("KthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKthSmallestIterative(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		k    int
		want int
	}{
		{
			name: "官方示例1",
			root: buildTree([]int{3, 1, 4, math.MinInt32, 2}),
			k:    1,
			want: 1,
		},
		{
			name: "官方示例2",
			root: buildTree([]int{5, 3, 6, 2, 4, math.MinInt32, math.MinInt32, 1}),
			k:    3,
			want: 3,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			k:    1,
			want: 1,
		},
		{
			name: "三个节点-找最大",
			root: buildTree([]int{2, 1, 3}),
			k:    3,
			want: 3,
		},
		{
			name: "三个节点-找中间",
			root: buildTree([]int{2, 1, 3}),
			k:    2,
			want: 2,
		},
		{
			name: "左斜树",
			root: buildTree([]int{3, 2, math.MinInt32, 1}),
			k:    2,
			want: 2,
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			k:    3,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KthSmallestIterative(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("KthSmallestIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkKthSmallest(b *testing.B) {
	root := buildTree([]int{5, 3, 6, 2, 4, math.MinInt32, math.MinInt32, 1})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KthSmallest(root, 3)
	}
}

func BenchmarkKthSmallestIterative(b *testing.B) {
	root := buildTree([]int{5, 3, 6, 2, 4, math.MinInt32, math.MinInt32, 1})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KthSmallestIterative(root, 3)
	}
}
