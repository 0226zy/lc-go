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
		{"示例1: 对称树", []int{1, 2, 2, 3, 4, 4, 3}, true},
		{"示例2: 非对称树", []int{1, 2, 2, nilNode, 3, nilNode, 3}, false},

		// 边界情况
		{"空树", nil, true},
		{"单节点", []int{1}, true},
		{"两节点相同值", []int{1, 2, 2}, true},
		{"两节点不同值", []int{1, 2, 3}, false},
		{"值相同且互为镜像", []int{1, 2, 2, nilNode, 3, 3}, true},
		{"值相同但不对称", []int{1, 2, 2, nilNode, 3, nilNode, 3, 4, 4}, false},
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
		{"节点数7(对称)", []int{1, 2, 2, 3, 4, 4, 3}},
		{"节点数63(对称)", buildSymmetricTreeVals(6)},
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

// buildSymmetricTreeVals 生成深度为 depth 的对称满二叉树的层序序列
// 同一层的节点值相同，保证任意镜像位置值相等
func buildSymmetricTreeVals(depth int) []int {
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
