package closestbinarysearchtreevalueii

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestClosestKValues(t *testing.T) {
	tests := []struct {
		name   string
		vals   []int
		target float64
		k      int
		want   []int
	}{
		// LeetCode 官方示例
		{"示例1: 平衡BST取2个", []int{4, 2, 5, 1, 3}, 3.714286, 2, []int{4, 3}},
		{"示例2: 单节点", []int{1}, 0.0, 1, []int{1}},

		// 边界：k 等于节点总数
		{"k等于节点总数", []int{4, 2, 5, 1, 3}, 100.0, 5, []int{1, 2, 3, 4, 5}},
		{"k等于1", []int{4, 2, 5, 1, 3}, 3.714286, 1, []int{4}},

		// 典型场景
		{"目标恰等于某节点值", []int{4, 2, 5, 1, 3}, 2.3, 2, []int{2, 3}},
		{"目标小于所有节点值", []int{4, 2, 5, 1, 3}, -10.0, 2, []int{1, 2}},
		{"目标大于所有节点值", []int{4, 2, 5, 1, 3}, 100.0, 2, []int{5, 4}},
		{"窗口在最左端", []int{10, 5, 15, 3, 7, 12, 20}, 0.5, 3, []int{3, 5, 7}},
		{"窗口在最右端", []int{10, 5, 15, 3, 7, 12, 20}, 99.0, 3, []int{20, 15, 12}},

		// 边界：退化树（链表形状）
		{"右斜树", []int{1, math.MinInt32, 2, math.MinInt32, 3, math.MinInt32, 4, math.MinInt32, 5}, 2.6, 3, []int{2, 3, 4}},
		{"左斜树", []int{5, 4, math.MinInt32, 3, math.MinInt32, 2, math.MinInt32, 1}, 4.4, 2, []int{4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := ClosestKValues(root, tt.target, tt.k)
			// 返回值顺序不限，使用无序比较
			if !utils.EqualIntSliceUnordered(got, tt.want) {
				t.Errorf("ClosestKValues(%v, %v, %d) = %v, want %v", tt.vals, tt.target, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkClosestKValues(b *testing.B) {
	// 构造 15 层满 BST（16383 个节点），层序填充即满足 BST 性质：
	// 每个节点的左子小于它、右子大于它
	depth := 15
	n := (1 << depth) - 1
	vals := make([]int, n)
	var fill func(idx, lo, hi int)
	fill = func(idx, lo, hi int) {
		if idx >= n || lo > hi {
			return
		}
		mid := lo + (hi-lo)/2
		vals[idx] = mid
		fill(2*idx+1, lo, mid-1)
		fill(2*idx+2, mid+1, hi)
	}
	fill(0, 1, n)
	root := datastructures.NewTreeFromSlice(vals)

	b.Run("满BST16383节点k100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ClosestKValues(root, 12345.6, 100)
		}
	})
}
