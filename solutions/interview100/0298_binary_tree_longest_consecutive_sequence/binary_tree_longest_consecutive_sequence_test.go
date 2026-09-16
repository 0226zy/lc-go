package binarytreelongestconsecutivesequence

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// null 层序切片中表示空节点的占位值（与 datastructures.NewTreeFromSlice 约定一致）
const null = math.MinInt32

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 最长路径 3->4->5", []int{1, null, 3, 2, 4, null, null, null, 5}, 3},
		{"示例2: 最长路径 2->3", []int{2, null, 3, 2, null, 1}, 2},

		// 边界：空树与单节点
		{"空树", []int{}, 0},
		{"单节点", []int{1}, 1},

		// 边界：整棵树是一条连续递增链
		{"全链连续递增", []int{1, null, 2, null, 3, null, 4}, 4},
		{"整棵树都连续", []int{1, 2, 3, 3, 4, 4, 5}, 3},

		// 边界：完全不连续
		{"全部相同值", []int{1, 1, 1}, 1},
		{"严格递减", []int{3, 2, 1}, 1},

		// 典型场景
		{"左子树更长", []int{5, 6, 2, 7, null, 3}, 3},
		{"根节点连续但右子树断开", []int{1, 2, 0, 3}, 3},
		{"含负数的连续序列", []int{-3, -2, null, -1}, 3},
		{"最长序列不在根出发", []int{1, 5, null, 6, null, 7, null, 8}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := LongestConsecutive(root); got != tt.want {
				t.Errorf("LongestConsecutive(%v) = %d, want %d", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkLongestConsecutive(b *testing.B) {
	// 构造一条长度为 10000 的连续递增链（全右子树）
	vals := make([]int, 0, 19999)
	for i := 0; i < 10000; i++ {
		vals = append(vals, i, null)
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestConsecutive(root)
	}
}
