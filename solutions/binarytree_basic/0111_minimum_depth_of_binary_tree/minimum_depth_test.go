package mindepth

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序切片，math.MinInt32 表示 nil
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [3,9,20,null,null,15,7]",
			[]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, 2},
		{"示例2: [2,null,3,null,4,null,5,null,6] 右斜链",
			[]int{2, math.MinInt32, 3, math.MinInt32, 4, math.MinInt32, 5, math.MinInt32, 6}, 5},

		// 边界：空树
		{"空树", nil, 0},
		// 边界：单节点
		{"单节点", []int{1}, 1},
		// 边界：左斜链（单侧子树为空，不能取 0）
		{"左斜链 [1,2,null,3,null,4]",
			[]int{1, 2, math.MinInt32, 3, math.MinInt32, 4}, 4},
		// 边界：根只有右子树且右子树很深
		{"根只有右子树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, 3},
		// 边界：完全二叉树，两侧均衡
		{"完全二叉树三层", []int{1, 2, 3, 4, 5, 6, 7}, 3},
		// 单侧子树为空的陷阱：根有左孩子（叶子）和右孩子（非叶子），最小深度是 2
		{"一侧是叶子一侧很深", []int{1, 2, 3, math.MinInt32, math.MinInt32, 4}, 2},
		// 含负数节点值，验证不受值影响
		{"含负数 [-1,-2,-3]", []int{-1, -2, -3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := MinDepth(root); got != tt.want {
				t.Errorf("MinDepth(%v) = %d, want %d", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkMinDepth(b *testing.B) {
	// 构造一条长度为 1000 的右斜链作为压力数据
	vals := make([]int, 0, 2000)
	for i := 1; i <= 1000; i++ {
		vals = append(vals, i)
		if i < 1000 {
			vals = append(vals, math.MinInt32)
		}
	}
	root := datastructures.NewTreeFromSlice(vals)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinDepth(root)
	}
}
