package largestbstsubtree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// nilNode 表示层序切片中的空节点
const nilNode = math.MinInt32

func TestLargestBSTSubtree(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: 最大BST子树含3个节点",
			[]int{10, 5, 15, 1, 8, nilNode, 7},
			3,
		},
		{
			"示例2: 最大BST子树含2个节点",
			[]int{4, 2, 7, 2, 3, 5, nilNode, 2, nilNode, nilNode, nilNode, nilNode, nilNode, 1},
			2,
		},

		// 边界：空树
		{
			"空树",
			nil,
			0,
		},

		// 边界：单节点
		{
			"单节点树",
			[]int{1},
			1,
		},

		// 典型场景：整棵树就是 BST
		{
			"整棵树都是BST",
			[]int{10, 5, 15, 1, 8, nilNode, 20},
			6,
		},

		// 典型场景：链状树
		{
			"单调递增右链整树是BST",
			[]int{1, nilNode, 2, nilNode, 3, nilNode, 4},
			4,
		},
		{
			"单调递减左链整树是BST",
			[]int{4, 3, nilNode, 2, nilNode, 1},
			4,
		},

		// 典型场景：只有叶子满足 BST
		{
			"根节点违反BST性质",
			[]int{5, 6, 4},
			1, // 根 5 的左孩子 6 > 5，只有单节点子树是 BST
		},

		// 典型场景：深层隐藏着较大的 BST
		{
			"深层隐藏较大BST",
			[]int{10, 8, 20, 6, 9, 15, 25, nilNode, nilNode, nilNode, nilNode, 5, nilNode},
			4, // 以 20 为根的子树 {20,15,25,5} 是 BST（5<15<20）；根 10 的右子树含 5 < 10，整树不是 BST
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := LargestBSTSubtree(root); got != tt.want {
				t.Errorf("LargestBSTSubtree(%v) = %d, want %d", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkLargestBSTSubtree(b *testing.B) {
	// 构造一棵 10 层满二叉树（1023 个节点），层序值即普通二叉树
	vals := make([]int, 0, 1023)
	for i := 1; i <= 1023; i++ {
		vals = append(vals, i)
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LargestBSTSubtree(root)
	}
}
