package binarytreelongestconsecutivesequenceii

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// null 是 NewTreeFromSlice 中表示空节点的占位值
const null = math.MinInt32

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3]", []int{1, 2, 3}, 2},
		{"示例2: [2,1,3]", []int{2, 1, 3}, 3},

		// 边界：空树与单节点
		{"空树", nil, 0},
		{"单节点", []int{1}, 1},

		// 边界：没有任何连续关系
		{"全不连续", []int{1, 3, 5}, 1},

		// 左链严格递减：3->2->1，整链即答案
		{"左链递减", []int{3, 2, null, 1}, 3},

		// 先向上再向下：路径 3->2->1->0，长度 4
		// 树形：
		//     1
		//    / \
		//   2   0
		//  /
		// 3
		{"子父子折线", []int{1, 2, 0, 3}, 4},

		// 右臂接递减、左臂接递增，拐点在根：
		//     2
		//    / \
		//   1   3
		// 路径 1->2->3 长度 3
		{"根为拐点", []int{2, 1, 3}, 3},

		// 负值也按相差 1 处理：-1->0->1
		{"包含负值", []int{0, -1, 1}, 3},
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
	// 手工构造一条 10000 层的连续左链：10000 -> 9999 -> ... -> 1
	root := &datastructures.TreeNode{Val: 10000}
	cur := root
	for v := 9999; v >= 1; v-- {
		cur.Left = &datastructures.TreeNode{Val: v}
		cur = cur.Left
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestConsecutive(root)
	}
}
