package binarytreemaximumpathsum

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3]", []int{1, 2, 3}, 6},
		{"示例2: 最优路径不经过根", []int{-10, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, 42},
		// 边界：单节点，答案即节点本身（即使为负）
		{"单节点", []int{1}, 1},
		{"单负节点", []int{-3}, -3},
		// 边界：全负二叉树，答案为最大的那个节点
		{"全负二叉树", []int{-1, -2, -3}, -1},
		// 边界：负贡献子树必须被舍弃
		{"舍弃负子树", []int{2, -1, -2}, 2},
		// 边界：单边链
		{"只有左子树", []int{1, 2, math.MinInt32, 3}, 6},
		// 边界：路径需要拐弯（同时取左右），最优路径为 8 -> 15 -> 20 -> 7 = 50
		{"路径拐弯", []int{-10, 9, 20, math.MinInt32, math.MinInt32, 15, 7, 8}, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := MaxPathSum(root); got != tt.want {
				t.Errorf("MaxPathSum(%v) = %d, want %d", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxPathSum(b *testing.B) {
	// 构建一个每层节点数翻倍的完全二叉树，规模约 4095 个节点，值在 [-100, 100] 间交替
	var vals []int
	for size := 1; len(vals)+size <= 4095; size *= 2 {
		for i := 0; i < size; i++ {
			vals = append(vals, (i%201)-100)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := datastructures.NewTreeFromSlice(vals)
		MaxPathSum(root)
	}
}
