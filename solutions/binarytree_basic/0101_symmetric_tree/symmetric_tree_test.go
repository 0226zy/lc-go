package symmetrictree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestIsSymmetric(t *testing.T) {
	nilNode := math.MinInt32
	tests := []struct {
		name string
		vals []int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: 对称的满二叉树", []int{1, 2, 2, 3, 4, 4, 3}, true},
		{"示例2: 结构不对称", []int{1, 2, 2, nilNode, 3, nilNode, 3}, false},

		// 边界情况
		{"空树", nil, true},
		{"单节点树", []int{1}, true},
		{"两子节点值相同", []int{1, 2, 2}, true},
		{"两子节点值不同", []int{1, 2, 3}, false},
		{"只有左孩子的链式树", []int{1, 2, nilNode, 3}, false},
		{"只有右孩子的链式树", []int{1, nilNode, 2, nilNode, 3}, false},
		{"含负数值且对称", []int{1, -2, -2, -3, -4, -4, -3}, true},
		{"含负数值且值不对称", []int{1, -2, -3}, false},
		{"值相同但结构不对称", []int{1, 2, 2, 2, nilNode, 2}, false},
		{"结构对称但值不对称", []int{1, 2, 2, 3, 4, 4, 5}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			if got := IsSymmetric(root); got != tt.want {
				t.Errorf("IsSymmetric(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkIsSymmetric(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"小规模对称树(7节点)", []int{1, 2, 2, 3, 4, 4, 3}},
		{"大规模对称满二叉树(1023节点)", buildSymmetricVals(10)},
		{"不对称树提前退出", append([]int{1, 2, 2}, math.MinInt32)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := datastructures.NewTreeFromSlice(bm.vals)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				IsSymmetric(root)
			}
		})
	}
}

// buildSymmetricVals 生成深度为 depth 的对称满二叉树的层序序列：
// 同一层所有节点取值相同（取层号），保证任意镜像位置的值必然相等。
func buildSymmetricVals(depth int) []int {
	n := 1 << depth
	vals := make([]int, n-1)
	for i := range vals {
		level := 0
		for x := i + 1; x > 1; x >>= 1 {
			level++
		}
		vals[i] = level
	}
	return vals
}
