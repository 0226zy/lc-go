package postordertraversal

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

// 解法列表，便于对每种实现跑同一套用例
var solutions = []struct {
	name string
	fn   func(*TreeNode) []int
}{
	{"递归法", PostorderTraversal},
	{"迭代法", PostorderTraversalIterative},
}

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序表示，math.MinInt32 表示空节点
		want []int
	}{
		{
			name: "官方示例1_根1右子2左孙3",
			vals: []int{1, math.MinInt32, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			name: "官方示例2_复杂树",
			vals: []int{1, 2, 3, 4, 5, math.MinInt32, 8, math.MinInt32, math.MinInt32, 6, 7, 9},
			want: []int{4, 6, 7, 5, 2, 9, 8, 3, 1},
		},
		{
			name: "空树返回空切片",
			vals: nil,
			want: []int{},
		},
		{
			name: "单节点树",
			vals: []int{1},
			want: []int{1},
		},
		{
			name: "左斜链式树",
			vals: []int{3, 2, math.MinInt32, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "右斜链式树",
			vals: []int{1, math.MinInt32, 2, math.MinInt32, 3},
			want: []int{3, 2, 1},
		},
		{
			name: "含负数的完全二叉树",
			vals: []int{-1, -2, 3, -4, 5, -6, 7},
			want: []int{-4, 5, -2, -6, 7, 3, -1},
		},
	}

	for _, tt := range tests {
		for _, sol := range solutions {
			t.Run(tt.name+"/"+sol.name, func(t *testing.T) {
				root := datastructures.NewTreeFromSlice(tt.vals)
				got := sol.fn(root)
				if got == nil {
					t.Fatalf("返回了 nil 切片，期望空切片 []")
				}
				if !utils.EqualIntSlice(got, tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func BenchmarkPostorderTraversal(b *testing.B) {
	// 构造一棵 4 层的满二叉树
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	root := datastructures.NewTreeFromSlice(vals)
	b.Run("递归法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PostorderTraversal(root)
		}
	})
	b.Run("迭代法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PostorderTraversalIterative(root)
		}
	})
}
