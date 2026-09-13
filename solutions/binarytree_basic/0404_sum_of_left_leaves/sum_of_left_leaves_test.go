package sumofleftleaves

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestSumOfLeftLeaves(t *testing.T) {
	nilNode := math.MinInt32
	tests := []struct {
		name string
		vals []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [3,9,20,null,null,15,7]", []int{3, 9, 20, nilNode, nilNode, 15, 7}, 24},
		{"示例2: 单节点树 [1]", []int{1}, 0},

		// 边界情况
		{"空树返回0", nil, 0},
		{"只有左孩子的两节点树", []int{1, 2}, 2},
		{"只有右孩子的两节点树", []int{1, nilNode, 2}, 0},
		{"纯左链树", []int{1, 2, nilNode, 3, nilNode, 4}, 4},
		{"纯右链树无左叶子", []int{1, nilNode, 2, nilNode, 3}, 0},
		{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, 4 + 6},
		{"含负数的左叶子", []int{1, -2, 3, -4}, -4},
		{"左孩子非叶子不计入", []int{1, 2, nilNode, 3, nilNode, nilNode, nilNode}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := SumOfLeftLeaves(root); got != tt.want {
				t.Errorf("SumOfLeftLeaves(%v) = %d, 期望 %d", tt.vals, got, tt.want)
			}
		})
	}
}

// makeChain 构造深度为 depth 的纯左链树的层序切片
func makeChain(depth int) []int {
	vals := make([]int, 0, 2*depth-1)
	for i := 1; i <= depth; i++ {
		vals = append(vals, i)
		if i < depth {
			vals = append(vals, math.MinInt32)
		}
	}
	return vals
}

// makeFullTree 构造层数为 levels 的满二叉树的层序切片
func makeFullTree(levels int) []int {
	n := 1<<levels - 1
	vals := make([]int, n)
	for i := range vals {
		vals[i] = i + 1
	}
	return vals
}

func BenchmarkSumOfLeftLeaves(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"链式树深度10", makeChain(10)},
		{"链式树深度100", makeChain(100)},
		{"链式树深度1000", makeChain(1000)},
		{"满二叉树10层", makeFullTree(10)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := datastructures.NewTreeFromSlice(bm.vals)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				SumOfLeftLeaves(root)
			}
		})
	}
}
