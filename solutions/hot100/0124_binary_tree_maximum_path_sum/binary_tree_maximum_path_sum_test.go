package binarytreemaximumpathsum

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want int
	}{
		{
			name: "官方示例1",
			root: buildTree([]int{1, 2, 3}),
			want: 6, // 2->1->3
		},
		{
			name: "官方示例2",
			root: buildTree([]int{-10, 9, 20, math.MinInt32, math.MinInt32, 15, 7}),
			want: 42, // 15->20->7
		},
		{
			name: "单节点为正",
			root: buildTree([]int{1}),
			want: 1,
		},
		{
			name: "单节点为负",
			root: buildTree([]int{-5}),
			want: -5,
		},
		{
			name: "全负数",
			root: buildTree([]int{-1, -2, -3}),
			want: -1, // 只能选一个节点
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: 18, // 1+max(7,10)=18 或 2+4+5=11 或 3+6+7=16
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: 6, // 1->2->3
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: 6, // 1->2->3
		},
		{
			name: "含负数和正数混合",
			root: buildTree([]int{-10, 9, 20, math.MinInt32, math.MinInt32, 15, 7}),
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxPathSum(tt.root)
			if got != tt.want {
				t.Errorf("MaxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMaxPathSum(b *testing.B) {
	root := buildTree([]int{-10, 9, 20, math.MinInt32, math.MinInt32, 15, 7})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxPathSum(root)
	}
}
