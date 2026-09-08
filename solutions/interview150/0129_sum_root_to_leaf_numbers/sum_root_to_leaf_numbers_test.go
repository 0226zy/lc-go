package sumroottoleafnumbers

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
		// 边界：单节点
		{"单节点", []int{7}, 7},
		// 边界：空树
		{"空树", []int{}, 0},
		// 边界：只有左子树的链 1->2->3 拼成 123
		{"只有左子树", []int{1, 2, math.MinInt32, 3}, 123},
		// 边界：只有右子树的链 1->2->3 拼成 123
		{"只有右子树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, 123},
		// 边界：含 0 的路径，0 在中间位而非前导位
		{"含0路径", []int{1, 0}, 10},
		// 边界：多位数字与多条路径累加
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
	// 构建一个每层节点数翻倍的完全二叉树，规模约 4095 个节点，值全为 1
	var vals []int
	for size := 1; len(vals)+size <= 4095; size *= 2 {
		for i := 0; i < size; i++ {
			vals = append(vals, 1)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := datastructures.NewTreeFromSlice(vals)
		SumNumbers(root)
	}
}
