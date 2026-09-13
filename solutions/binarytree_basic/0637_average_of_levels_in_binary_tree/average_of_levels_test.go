package averageoflevels

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

const eps = 1e-5

// equalFloatSlice 按 eps 精度比较两个浮点切片；均视为空时返回 true
func equalFloatSlice(got, want []float64) bool {
	if len(got) == 0 && len(want) == 0 {
		return true
	}
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if math.Abs(got[i]-want[i]) > eps {
			return false
		}
	}
	return true
}

func TestAverageOfLevels(t *testing.T) {
	tests := []struct {
		name string
		vals []int // math.MinInt32 表示空节点
		want []float64
	}{
		// LeetCode 官方示例
		{"示例1_完整层", []int{3, 9, 20, 15, 7}, []float64{3.0, 14.5, 11.0}},
		{"示例2_含空节点", []int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, []float64{3.0, 14.5, 11.0}},

		// 边界情况
		{"空树返回nil", nil, nil},
		{"单节点", []int{42}, []float64{42.0}},
		{"左链式树", []int{1, 2, math.MinInt32, 3, math.MinInt32, 4}, []float64{1.0, 2.0, 3.0, 4.0}},
		{"含负数平均", []int{-1, -2, -3}, []float64{-1.0, -2.5}},
		{"正负混合平均为零", []int{5, -5, 5, -5, 5}, []float64{5.0, 0.0, 0.0}},
		{"平均恰为整数", []int{4, 2, 6}, []float64{4.0, 4.0}},
		{"不平衡树", []int{1, 2, math.MinInt32, 3, 4}, []float64{1.0, 2.0, 3.5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := AverageOfLevels(root)
			if !equalFloatSlice(got, tt.want) {
				t.Errorf("AverageOfLevels(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkAverageOfLevels(b *testing.B) {
	// 深度 12 的满二叉树，共 4095 个节点
	depth := 12
	counter := 1
	var build func(d int) *datastructures.TreeNode
	build = func(d int) *datastructures.TreeNode {
		if d == 0 {
			return nil
		}
		node := &datastructures.TreeNode{Val: counter}
		counter++
		node.Left = build(d - 1)
		node.Right = build(d - 1)
		return node
	}
	root := build(depth)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AverageOfLevels(root)
	}
}
