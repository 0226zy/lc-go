package isbalanced

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: [3,9,20,null,null,15,7] 是平衡树",
			[]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, true},
		{"示例2: [1,2,2,3,3,null,null,4,4] 不是平衡树",
			[]int{1, 2, 2, 3, 3, math.MinInt32, math.MinInt32, 4, 4}, false},
		{"示例3: 空树 是平衡树", nil, true},

		// 边界情况
		{"单节点树", []int{1}, true},
		{"两个节点的左斜树", []int{1, 2}, true},
		{"两个节点的右斜树", []int{1, math.MinInt32, 2}, true},
		{"链式树(全左)", []int{1, 2, math.MinInt32, 3, math.MinInt32, 4}, false},
		{"链式树(全右)",
			[]int{1, math.MinInt32, 2, math.MinInt32, 3, math.MinInt32, 4}, false},
		{"完全二叉树", []int{1, 2, 3, 4, 5, 6, 7}, true},
		{"含负数的平衡树", []int{-1, -2, -3, -4, -5, -6}, true},
		{"根节点左右高度差恰为1", []int{1, 2, 3, 4, 5}, true},
		{"深层子树不平衡(提前剪枝场景)",
			[]int{1, 2, 2, 3, math.MinInt32, math.MinInt32, 3, 4, math.MinInt32, math.MinInt32, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := IsBalanced(root); got != tt.want {
				t.Errorf("IsBalanced(%v) = %v, want %v\n树结构:\n%s", tt.vals, got, tt.want, root)
			}
		})
	}
}

func BenchmarkIsBalanced(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"满二叉树-深度10", buildFullTreeVals(10)},
		{"不平衡链式树-长度1000", buildChainVals(1000)},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := datastructures.NewTreeFromSlice(bm.vals)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				IsBalanced(root)
			}
		})
	}
}

// buildFullTreeVals 构造深度为 depth 的满二叉树的层序切片（共 2^depth-1 个节点）
func buildFullTreeVals(depth int) []int {
	n := 1<<depth - 1
	vals := make([]int, n)
	for i := range vals {
		vals[i] = i + 1
	}
	return vals
}

// buildChainVals 构造长度为 n 的左链树的层序切片
func buildChainVals(n int) []int {
	vals := make([]int, 2*n-1)
	for i := 0; i < n; i++ {
		vals[2*i] = i + 1
		if 2*i+1 < len(vals) {
			vals[2*i+1] = math.MinInt32
		}
	}
	return vals
}
