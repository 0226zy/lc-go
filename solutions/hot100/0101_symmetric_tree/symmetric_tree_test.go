package symmetrictree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want bool
	}{
		{
			name: "空树",
			root: nil,
			want: true,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: true,
		},
		{
			name: "官方示例1-对称",
			root: buildTree([]int{1, 2, 2, 3, 4, 4, 3}),
			want: true,
		},
		{
			name: "官方示例2-不对称",
			root: buildTree([]int{1, 2, 2, math.MinInt32, 3, math.MinInt32, 3}),
			want: false,
		},
		{
			name: "左斜树-不对称",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: false,
		},
		{
			name: "只有根和两个相同子节点-对称",
			root: buildTree([]int{1, 2, 2}),
			want: true,
		},
		{
			name: "根的左右子节点值不同-不对称",
			root: buildTree([]int{1, 2, 3}),
			want: false,
		},
		{
			name: "三层完全对称",
			root: buildTree([]int{1, 2, 2, 3, 3, 3, 3}),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSymmetric(tt.root)
			if got != tt.want {
				t.Errorf("IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSymmetricIterative(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want bool
	}{
		{
			name: "空树",
			root: nil,
			want: true,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: true,
		},
		{
			name: "官方示例1-对称",
			root: buildTree([]int{1, 2, 2, 3, 4, 4, 3}),
			want: true,
		},
		{
			name: "官方示例2-不对称",
			root: buildTree([]int{1, 2, 2, math.MinInt32, 3, math.MinInt32, 3}),
			want: false,
		},
		{
			name: "左斜树-不对称",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: false,
		},
		{
			name: "只有根和两个相同子节点-对称",
			root: buildTree([]int{1, 2, 2}),
			want: true,
		},
		{
			name: "根的左右子节点值不同-不对称",
			root: buildTree([]int{1, 2, 3}),
			want: false,
		},
		{
			name: "三层完全对称",
			root: buildTree([]int{1, 2, 2, 3, 3, 3, 3}),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSymmetricIterative(tt.root)
			if got != tt.want {
				t.Errorf("IsSymmetricIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsSymmetric(b *testing.B) {
	root := buildTree([]int{1, 2, 2, 3, 4, 4, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsSymmetric(root)
	}
}

func BenchmarkIsSymmetricIterative(b *testing.B) {
	root := buildTree([]int{1, 2, 2, 3, 4, 4, 3})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsSymmetricIterative(root)
	}
}
