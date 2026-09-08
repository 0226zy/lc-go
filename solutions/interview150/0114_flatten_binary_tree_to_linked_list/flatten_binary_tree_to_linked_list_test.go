package flattenbinarytreetolinkedlist

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// rightChainToSlice 沿着 Right 指针收集链表值
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
		// 边界：只有左子树（链表向右下拐两次）
		{"只有左子树", []int{1, 2, math.MinInt32, 3}, []int{1, 2, 3}},
		// 边界：只有右子树（本身已是链表）
		{"只有右子树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, []int{1, 2, 3}},
		// 边界：右子树挂到左子树最右节点之下
		{"右子树深于左子树", []int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1, 2, 4, 8, 5, 3, 6, 7}},
		// 边界：左子树深于右子树
		{"左子树深于右子树", []int{1, 2, 3, 4, 5, 6, 7, math.MinInt32, math.MinInt32, 8}, []int{1, 2, 4, 5, 8, 3, 6, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			Flatten(root)
			// 先检查所有 Left 指针已置空
			for curr := root; curr != nil; curr = curr.Right {
				if curr.Left != nil {
					t.Fatalf("展开后节点 %d 的 Left 指针未置空", curr.Val)
				}
			}
			got := rightChainToSlice(root)
			if tt.want == nil {
				if got != nil {
					t.Errorf("Flatten(%v) 的右链表 = %v, want nil", tt.vals, got)
				}
				return
			}
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten(%v) 的右链表 = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkFlatten(b *testing.B) {
	// 构建一个每层节点数翻倍的完全二叉树，规模约 4095 个节点
	var vals []int
	for size := 1; len(vals)+size <= 4095; size *= 2 {
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
