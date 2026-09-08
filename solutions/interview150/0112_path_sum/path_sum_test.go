package pathsum

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		vals      []int
		targetSum int
		want      bool
	}{
		// LeetCode 官方示例
		{"示例1: 存在路径和为22", []int{5, 4, 8, 11, math.MinInt32, 13, 4, 7, 2, math.MinInt32, math.MinInt32, math.MinInt32, 1}, 22, true},
		{"示例2: 不存在路径和为5", []int{1, 2, 3}, 5, false},
		{"示例3: 空树", []int{}, 0, false},
		// 边界：单节点，值恰好匹配 / 不匹配
		{"单节点值匹配", []int{1}, 1, true},
		{"单节点值不匹配", []int{1}, 2, false},
		// 边界：负数
		{"负数路径和", []int{-2, math.MinInt32, -3}, -5, true},
		{"负数但无匹配", []int{-2, math.MinInt32, -3}, -2, false},
		// 边界：目标和为 0，需存在真实路径 1 -> -1
		{"目标和为0", []int{1, -1}, 0, true},
		// 边界：targetSum 为负，路径含正负抵消
		{"正负抵消", []int{1, -2, math.MinInt32, 3}, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := HasPathSum(root, tt.targetSum); got != tt.want {
				t.Errorf("HasPathSum(%v, %d) = %v, want %v", tt.vals, tt.targetSum, got, tt.want)
			}
		})
	}
}

func BenchmarkHasPathSum(b *testing.B) {
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
		// 选一个不可能成立的目标和，迫使遍历整棵树
		HasPathSum(root, math.MaxInt32)
	}
}
