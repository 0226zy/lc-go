package twosumbsts

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// 层序切片中的空位用 math.MinInt32 表示（见 datastructures.NewTreeFromSlice）
const null = math.MinInt32

func TestTwoSumBSTs(t *testing.T) {
	tests := []struct {
		name   string
		root1  []int
		root2  []int
		target int
		want   bool
	}{
		// LeetCode 官方示例
		{"示例1: 2+3=5", []int{2, 1, 4}, []int{1, 0, 3}, 5, true},
		{"示例2: 无法凑出18", []int{0, -10, 10}, []int{5, 1, 7, 0, 2}, 18, false},

		// 边界：单节点树
		{"单节点恰好匹配", []int{1}, []int{2}, 3, true},
		{"单节点不匹配", []int{1}, []int{2}, 4, false},

		// 边界：负数与极大值
		{"负数相加", []int{-5, -10, 0}, []int{5}, 0, true},
		{"负数凑负目标", []int{-10}, []int{-20, -30}, -40, true},
		{"极大值相加", []int{1000000000}, []int{-1000000000}, 0, true},

		// 边界：同一棵树中的两个值不能凑（必须分别来自两棵树）
		{"需要跨树配对", []int{1, null, 2}, []int{3}, 5, true},
		{"目标只能由同树节点凑出", []int{1, 5}, []int{2}, 6, false},

		// 典型场景：深层子树中命中
		{"深层节点命中", []int{4, 2, 6, 1, 3, 5, 7}, []int{8, null, 9, null, null, null, 10}, 11, true},
		{"完全无解", []int{1, 2, 3}, []int{10, 20}, 100, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root1 := datastructures.NewTreeFromSlice(tt.root1)
			root2 := datastructures.NewTreeFromSlice(tt.root2)
			if got := TwoSumBSTs(root1, root2, tt.target); got != tt.want {
				t.Errorf("TwoSumBSTs(%v, %v, %d) = %v, want %v", tt.root1, tt.root2, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkTwoSumBSTs(b *testing.B) {
	// 构造一棵约 5000 节点的左斜树和一棵右斜树
	vals1 := make([]int, 5000)
	vals2 := make([]int, 5000)
	for i := range vals1 {
		vals1[i] = i
		vals2[i] = -i
	}
	root1 := buildSkewed(vals1)
	root2 := buildSkewed(vals2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSumBSTs(root1, root2, 4999)
	}
}

// buildSkewed 用切片值构造一条右斜树，避免测试中出现极深层序切片
func buildSkewed(vals []int) *datastructures.TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &datastructures.TreeNode{Val: vals[0]}
	cur := root
	for _, v := range vals[1:] {
		cur.Right = &datastructures.TreeNode{Val: v}
		cur = cur.Right
	}
	return root
}
