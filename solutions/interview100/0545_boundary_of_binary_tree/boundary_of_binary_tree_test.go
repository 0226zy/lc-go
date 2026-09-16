package boundaryofbinarytree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

// null 是 NewTreeFromSlice 中表示空节点的占位值
const null = math.MinInt32

func TestBoundaryOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,4]", []int{1, 2, 3, 4}, []int{1, 2, 4, 3}},
		{"示例2: [1,null,2,3,4]", []int{1, null, 2, 3, 4}, []int{1, 3, 4, 2}},

		// 边界：空树与单节点
		{"空树", nil, nil},
		{"根即叶子", []int{1}, []int{1}},

		// 边界：只有左链（右边界为空）
		{"只有左链", []int{1, 2, null, 3}, []int{1, 2, 3}},

		// 边界：只有右链（左边界为空，右边界逆序输出）
		{"只有右链", []int{1, null, 2, null, 3}, []int{1, 3, 2}},

		// 边界：根只有左孩子且左孩子是叶子
		{"根只有左叶子", []int{1, 2}, []int{1, 2}},

		// 典型场景：满二叉树
		{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 4, 5, 6, 7, 3}},

		// 典型场景：左右边界中间隔着多层叶子
		{"复杂树", []int{1, 2, 3, 4, null, 5, null, 6, 7, 8, 9},
			[]int{1, 2, 4, 6, 7, 8, 9, 5, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := BoundaryOfBinaryTree(root); !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("BoundaryOfBinaryTree(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkBoundaryOfBinaryTree(b *testing.B) {
	// 构造一棵 10 层的满二叉树（1023 个节点）
	vals := make([]int, 1023)
	for i := range vals {
		vals[i] = i + 1
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoundaryOfBinaryTree(root)
	}
}
