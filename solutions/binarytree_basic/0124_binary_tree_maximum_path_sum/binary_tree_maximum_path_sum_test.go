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
		{"示例1: 最优路径经过根 [1,2,3]", []int{1, 2, 3}, 6},
		{"示例2: 最优路径不经过根 [-10,9,20,null,null,15,7]", []int{-10, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, 42},
		// 边界情况
		{"空树返回0", nil, 0},
		{"单节点", []int{1}, 1},
		{"单节点为负数", []int{-3}, -3},
		{"全为负数的树取最大单节点", []int{-1, -2, -3}, -1},
		{"负贡献子树被舍弃", []int{2, -1, -2}, 2},
		// 链式树：只有左孩子，最优路径是整条链 3 -> 2 -> 1
		{"左链式树", []int{1, 2, math.MinInt32, 3}, 6},
		// 最优路径需要拐弯：9 -> -10 -> 20 -> 15 中 15 子树更优时为 8 -> 15 -> 20 -> 7 = 50
		{"路径在内部节点拐弯", []int{-10, 9, 20, math.MinInt32, math.MinInt32, 15, 7, 8}, 50},
		// 最大值恰好在最深的叶子出发的路径上
		{"深叶子出发的路径", []int{5, 4, 8, 11, math.MinInt32, 13, 4, 7, 2, math.MinInt32, math.MinInt32, math.MinInt32, 1}, 48},
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
	// 构造一棵 12 层完全二叉树（4095 个节点），节点值在 [-100, 100] 间交替
	var vals []int
	for size := 1; len(vals)+size <= 4095; size *= 2 {
		for i := 0; i < size; i++ {
			vals = append(vals, (i%201)-100)
		}
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxPathSum(root)
	}
}
