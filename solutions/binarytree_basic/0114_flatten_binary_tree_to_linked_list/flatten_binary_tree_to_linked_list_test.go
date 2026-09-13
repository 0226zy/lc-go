package flattenbinarytreetolinkedlist

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// rightChainToSlice 沿着 Right 指针收集展开后链表上的节点值
func rightChainToSlice(root *datastructures.TreeNode) []int {
	var result []int
	for curr := root; curr != nil; curr = curr.Right {
		result = append(result, curr.Val)
	}
	return result
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,5,3,4,null,6]", []int{1, 2, 5, 3, 4, math.MinInt32, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"示例2: 空树", []int{}, []int{}},
		{"示例3: 单节点", []int{0}, []int{0}},
		// 边界：只有左子树，展开后变成向右的链
		{"只有左子树", []int{1, 2, math.MinInt32, 3}, []int{1, 2, 3}},
		// 边界：只有右子树，本身就是目标形态
		{"只有右子树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, []int{1, 2, 3}},
		// 边界：左链与右链混合的链式树
		{"左右混合的链式树", []int{1, 2, math.MinInt32, math.MinInt32, 3, math.MinInt32, math.MinInt32, math.MinInt32, 4}, []int{1, 2, 3, 4}},
		// 边界：含负数节点值
		{"含负数节点", []int{-1, -2, 5, -3, math.MinInt32, math.MinInt32, 6}, []int{-1, -2, -3, 5, 6}},
		// 边界：满二叉树
		{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 4, 5, 3, 6, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			Flatten(root)
			// 先检查所有 Left 指针均已置空
			for curr := root; curr != nil; curr = curr.Right {
				if curr.Left != nil {
					t.Fatalf("展开后节点 %d 的 Left 指针未置空", curr.Val)
				}
			}
			got := rightChainToSlice(root)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten(%v) 展开后的链表 = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkFlatten(b *testing.B) {
	// 构建一棵约 2047 个节点的满二叉树（11 层）
	var vals []int
	for size := 1; len(vals)+size <= 2047; size *= 2 {
		for i := 0; i < size; i++ {
			vals = append(vals, i+1)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := datastructures.NewTreeFromSlice(vals)
		Flatten(root)
	}
}
