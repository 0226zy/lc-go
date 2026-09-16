package binarytreeverticalordertraversal

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

// nilNode 表示层序切片中的空节点
const nilNode = math.MinInt32

func TestVerticalOrder(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want [][]int
	}{
		// LeetCode 官方示例
		{
			"示例1: 普通二叉树",
			[]int{3, 9, 20, nilNode, nilNode, 15, 7},
			[][]int{{9}, {3, 15}, {20}, {7}},
		},
		{
			"示例2: 同列多节点按层序排列",
			[]int{3, 9, 8, 4, 0, 1, 7},
			[][]int{{4}, {9}, {3, 0, 1}, {8}, {7}},
		},
		{
			"示例3: 左右孩子缺失的同列场景",
			[]int{3, 9, 8, 4, 0, 1, 7, nilNode, nilNode, nilNode, 2, 5},
			[][]int{{4}, {9, 5}, {3, 0, 1}, {8, 2}, {7}},
		},

		// 边界：空树
		{
			"空树",
			nil,
			nil,
		},

		// 边界：单节点
		{
			"单节点树",
			[]int{1},
			[][]int{{1}},
		},

		// 边界：只有左子树（列号全部为负）
		{
			"只有左链",
			[]int{1, 2, nilNode, 3},
			[][]int{{3}, {2}, {1}},
		},

		// 边界：只有右子树（列号全部为正）
		{
			"只有右链",
			[]int{1, nilNode, 2, nilNode, 3},
			[][]int{{1}, {2}, {3}},
		},

		// 典型场景：完全二叉树
		{
			"完全二叉树",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := VerticalOrder(root)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("VerticalOrder(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkVerticalOrder(b *testing.B) {
	// 构造一棵 10 层满二叉树的层序切片（1023 个节点）
	vals := make([]int, 0, 1023)
	for i := 1; i <= 1023; i++ {
		vals = append(vals, i)
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		VerticalOrder(root)
	}
}
