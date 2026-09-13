package sumnumbers

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3]", []int{1, 2, 3}, 25},
		{"示例2: [4,9,0,5,1]", []int{4, 9, 0, 5, 1}, 1026},
		// 边界：空树
		{"空树", []int{}, 0},
		// 边界：单节点
		{"单节点", []int{7}, 7},
		// 边界：只有左子树的链 1->2->3 拼成 123
		{"只有左子树的链", []int{1, 2, math.MinInt32, 3}, 123},
		// 边界：只有右子树的链 1->2->3 拼成 123
		{"只有右子树的链", []int{1, math.MinInt32, 2, math.MinInt32, 3}, 123},
		// 边界：路径中含 0，0 在中间位而非前导位
		{"路径含0", []int{1, 0}, 10},
		// 边界：根节点值为 0，拼出的数字以 0 开头但数值不受影响
		{"根节点为0", []int{0, 1}, 1},
		// 边界：多条路径、多位数字累加，满二叉树 1->7
		{"多路径累加", []int{1, 2, 3, 4, 5, 6, 7}, 522},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := SumNumbers(root); got != tt.want {
				t.Errorf("SumNumbers(%v) = %d, want %d", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkSumNumbers(b *testing.B) {
	// 构建一棵 4095 个节点的满二叉树（12 层），节点值全为 1
	var vals []int
	for size := 1; len(vals)+size <= 4095; size *= 2 {
		for i := 0; i < size; i++ {
			vals = append(vals, 1)
		}
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumNumbers(root)
	}
}
