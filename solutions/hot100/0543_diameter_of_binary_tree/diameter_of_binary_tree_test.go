package diameterofbinarytree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want int
	}{
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: 0,
		},
		{
			name: "官方示例1",
			root: buildTree([]int{1, 2, 3, 4, 5}),
			want: 3,
		},
		{
			name: "官方示例2",
			root: buildTree([]int{1, 2}),
			want: 1,
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: 2,
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: 2,
		},
		{
			name: "完全二叉树三层",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: 4,
		},
		{
			name: "直径经过根节点-深层左子树",
			//       1
			//      / \
			//     2   3
			//    /
			//   4
			//  / \
			// 5   6
			// 最长路径: 5→4→2→1→3 或 6→4→2→1→3，共4条边
			root: buildTree([]int{1, 2, 3, 4, math.MinInt32, math.MinInt32, math.MinInt32, 5, 6}),
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DiameterOfBinaryTree(tt.root)
			if got != tt.want {
				t.Errorf("DiameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDiameterOfBinaryTree(b *testing.B) {
	root := buildTree([]int{1, 2, 3, 4, 5})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiameterOfBinaryTree(root)
	}
}
