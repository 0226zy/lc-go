package validatebinarysearchtree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want bool
	}{
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: true,
		},
		{
			name: "官方示例1-有效BST",
			root: buildTree([]int{2, 1, 3}),
			want: true,
		},
		{
			name: "官方示例2-无效BST",
			root: buildTree([]int{5, 1, 4, math.MinInt32, math.MinInt32, 3, 6}),
			want: false,
		},
		{
			name: "左子树节点值等于根节点-无效",
			root: buildTree([]int{1, 1}),
			want: false,
		},
		{
			name: "右子树节点值等于根节点-无效",
			root: buildTree([]int{1, math.MinInt32, 1}),
			want: false,
		},
		{
			name: "深层节点违反BST性质",
			//       10
			//      /  \
			//     5    15
			//         /  \
			//        6    20
			// 6 不应该在 15 的左子树中（6 < 10）
			root: buildTree([]int{10, 5, 15, math.MinInt32, math.MinInt32, 6, 20}),
			want: false,
		},
		{
			name: "完全二叉树BST",
			root: buildTree([]int{4, 2, 6, 1, 3, 5, 7}),
			want: true,
		},
		{
			name: "左斜树BST",
			root: buildTree([]int{3, 2, math.MinInt32, 1}),
			want: true,
		},
		{
			name: "节点值为int32边界",
			// 用 int32 最小值和最大值附近的值测试
			root: buildTree([]int{0, math.MinInt32, math.MaxInt32}),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidBST(tt.root)
			if got != tt.want {
				t.Errorf("IsValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidBSTInorder(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want bool
	}{
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: true,
		},
		{
			name: "官方示例1-有效BST",
			root: buildTree([]int{2, 1, 3}),
			want: true,
		},
		{
			name: "官方示例2-无效BST",
			root: buildTree([]int{5, 1, 4, math.MinInt32, math.MinInt32, 3, 6}),
			want: false,
		},
		{
			name: "左子树节点值等于根节点-无效",
			root: buildTree([]int{1, 1}),
			want: false,
		},
		{
			name: "右子树节点值等于根节点-无效",
			root: buildTree([]int{1, math.MinInt32, 1}),
			want: false,
		},
		{
			name: "深层节点违反BST性质",
			root: buildTree([]int{10, 5, 15, math.MinInt32, math.MinInt32, 6, 20}),
			want: false,
		},
		{
			name: "完全二叉树BST",
			root: buildTree([]int{4, 2, 6, 1, 3, 5, 7}),
			want: true,
		},
		{
			name: "左斜树BST",
			root: buildTree([]int{3, 2, math.MinInt32, 1}),
			want: true,
		},
		{
			name: "节点值为int32边界",
			root: buildTree([]int{0, math.MinInt32, math.MaxInt32}),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidBSTInorder(tt.root)
			if got != tt.want {
				t.Errorf("IsValidBSTInorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsValidBST(b *testing.B) {
	root := buildTree([]int{4, 2, 6, 1, 3, 5, 7})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsValidBST(root)
	}
}

func BenchmarkIsValidBSTInorder(b *testing.B) {
	root := buildTree([]int{4, 2, 6, 1, 3, 5, 7})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsValidBSTInorder(root)
	}
}
