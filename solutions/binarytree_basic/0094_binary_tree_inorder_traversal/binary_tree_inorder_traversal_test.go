package inordertraversal

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序切片，math.MinInt32 表示空节点
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: 树[1,null,2,3]",
			[]int{1, math.MinInt32, 2, 3},
			[]int{1, 3, 2}},
		{"示例2: 空树",
			nil,
			[]int{}},
		{"示例3: 单节点",
			[]int{1},
			[]int{1}},

		// 边界情况
		{"左斜链式树",
			[]int{3, 2, math.MinInt32, 1},
			[]int{1, 2, 3}},
		{"右斜链式树",
			[]int{1, math.MinInt32, 2, math.MinInt32, 3},
			[]int{1, 2, 3}},
		{"含负数节点",
			[]int{-1, -2, 3, math.MinInt32, -4},
			[]int{-2, -4, -1, 3}},
		{"完全二叉树",
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{4, 2, 5, 1, 6, 3, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := InorderTraversal(root)
			if !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
			gotIter := InorderTraversalIterative(root)
			if !utils.EqualIntSlice(gotIter, tt.want) {
				t.Errorf("InorderTraversalIterative() = %v, want %v", gotIter, tt.want)
			}
		})
	}
}

func BenchmarkInorderTraversal(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"完全二叉树7节点", []int{1, 2, 3, 4, 5, 6, 7}},
		{"左斜链式1000节点", buildLeftSkewed(1000)},
	}
	for _, bm := range benchmarks {
		root := datastructures.NewTreeFromSlice(bm.vals)
		b.Run(bm.name+"_递归", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				InorderTraversal(root)
			}
		})
		b.Run(bm.name+"_迭代", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				InorderTraversalIterative(root)
			}
		})
	}
}

// buildLeftSkewed 生成一棵 n 个节点的左斜链式树的层序切片
func buildLeftSkewed(n int) []int {
	vals := make([]int, 0, 2*n-1)
	vals = append(vals, n)
	for i := n - 1; i >= 1; i-- {
		vals = append(vals, i)             // 左子
		vals = append(vals, math.MinInt32) // 右子为空
	}
	return vals
}
