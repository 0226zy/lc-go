package maximumdepthofbinarytree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want int
	}{
		{
			name: "空树",
			root: nil,
			want: 0,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: 1,
		},
		{
			name: "官方示例1",
			root: buildTree([]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}),
			want: 3,
		},
		{
			name: "官方示例2",
			root: buildTree([]int{1, math.MinInt32, 2}),
			want: 2,
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: 3,
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: 3,
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxDepth(tt.root)
			if got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxDepthBFS(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want int
	}{
		{
			name: "空树",
			root: nil,
			want: 0,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: 1,
		},
		{
			name: "官方示例1",
			root: buildTree([]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}),
			want: 3,
		},
		{
			name: "官方示例2",
			root: buildTree([]int{1, math.MinInt32, 2}),
			want: 2,
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: 3,
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: 3,
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxDepthBFS(tt.root)
			if got != tt.want {
				t.Errorf("MaxDepthBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	root := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxDepth(root)
	}
}

func BenchmarkMaxDepthBFS(b *testing.B) {
	root := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxDepthBFS(root)
	}
}
