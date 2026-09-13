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
		// 边界：单节点树，值匹配 / 不匹配
		{"单节点值匹配", []int{1}, 1, true},
		{"单节点值不匹配", []int{1}, 2, false},
		// 边界：路径含负数节点
		{"负数路径和", []int{-2, math.MinInt32, -3}, -5, true},
		{"负数但无匹配路径", []int{-2, math.MinInt32, -3}, -2, false},
		// 边界：目标和为 0，必须依赖真实路径 1 -> -1 才能成立
		{"目标和为0存在路径", []int{1, -1}, 0, true},
		{"目标和为0不存在路径", []int{1, 2}, 0, false},
		// 边界：左子树失败但右子树成功，验证递归两分支
		{"右子树命中", []int{1, 2, 3}, 4, true},
		// 边界：链状树
		{"链状树匹配", []int{1, math.MinInt32, 2, math.MinInt32, 3}, 6, true},
		{"链状树不匹配", []int{1, math.MinInt32, 2, math.MinInt32, 3}, 5, false},
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
	// 构建一棵约 4095 个节点的满二叉树，节点值全为 1
	var vals []int
	for size := 1; len(vals)+size <= 4095; size *= 2 {
		for i := 0; i < size; i++ {
			vals = append(vals, 1)
		}
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 目标和取不可能成立的值，迫使遍历整棵树
		HasPathSum(root, math.MaxInt32)
	}
}
