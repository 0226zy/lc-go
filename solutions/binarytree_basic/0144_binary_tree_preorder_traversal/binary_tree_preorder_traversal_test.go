package preordertraversal

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

// solutions 存放本题的各个解法，方便统一做表驱动测试
var solutions = []struct {
	name string
	fn   func(*TreeNode) []int
}{
	{"递归法", PreorderTraversal},
	{"迭代法", PreorderTraversalIterative},
}

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序构造输入，math.MinInt32 表示空节点
		want []int
	}{
		{
			name: "官方示例一：右斜树 [1,null,2,3]",
			vals: []int{1, math.MinInt32, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "官方示例二：完整树 [1,2,3,4,5,null,8,null,null,6,7,9]",
			vals: []int{1, 2, 3, 4, 5, math.MinInt32, 8, math.MinInt32, math.MinInt32, 6, 7, 9},
			want: []int{1, 2, 4, 5, 6, 7, 3, 8, 9},
		},
		{
			name: "空树返回空切片",
			vals: []int{},
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
			want: []int{3, 2, 1},
		},
		{
			name: "含负数的完全二叉树",
			vals: []int{-1, -2, 3, -4, 5},
			want: []int{-1, -2, -4, 5, 3},
		},
	}

	for _, tt := range tests {
		root := datastructures.NewTreeFromSlice(tt.vals)
		for _, sol := range solutions {
			t.Run(tt.name+"/"+sol.name, func(t *testing.T) {
				got := sol.fn(root)
				if !utils.EqualIntSlice(got, tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	}
}

// TestPreorderTraversal_EmptyNotNil 空树时结果应为空切片而非 nil，方便调用方直接遍历
func TestPreorderTraversal_EmptyNotNil(t *testing.T) {
	for _, sol := range solutions {
		t.Run(sol.name, func(t *testing.T) {
			got := sol.fn(nil)
			if got == nil {
				t.Errorf("空树应返回空切片而非 nil")
			}
			if len(got) != 0 {
				t.Errorf("got %v, want []", got)
			}
		})
	}
}

func BenchmarkPreorderTraversal(b *testing.B) {
	// 构造一棵 1023 节点的满二叉树
	vals := make([]int, 1023)
	for i := range vals {
		vals[i] = i + 1
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.Run("递归法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PreorderTraversal(root)
		}
	})
	b.Run("迭代法", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PreorderTraversalIterative(root)
		}
	})
}
