package averageoflevelsinbinarytree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestAverageOfLevels(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []float64
	}{
		// LeetCode 官方示例
		{"示例1: [3,9,20,15,7]", []int{3, 9, 20, 15, 7}, []float64{3.0, 14.5, 11.0}},
		{"示例2: [3,9,20,null,null,15,7]", []int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, []float64{3.0, 14.5, 11.0}},

		// 边界：空树
		{"空树", nil, nil},
		// 边界：单节点
		{"单节点", []int{1}, []float64{1.0}},
		// 边界：负数与极值
		{"负数", []int{-1, -2, -3}, []float64{-1.0, -2.5}},
		// 边界：整除无余（平均值恰为整数）
		{"整除平均", []int{4, 2, 6}, []float64{4.0, 4.0}},
		// 边界：只有左子树
		{"只有左子树", []int{1, 2, math.MinInt32, 3}, []float64{1.0, 2.0, 3.0}},
		// 边界：较大值验证浮点求和正确性（题目约束 Node.val <= 10^5，float64 精度足够）
		{"大值求和", []int{1000000, 1000000, 1000000}, []float64{1000000.0, 1000000.0}},
	}

	const eps = 1e-5
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := AverageOfLevels(root)
			if tt.want == nil {
				if got != nil && len(got) != 0 {
					t.Errorf("AverageOfLevels(%v) = %v, want nil", tt.vals, got)
				}
				return
			}
			if len(got) != len(tt.want) {
				t.Fatalf("AverageOfLevels(%v) = %v, want %v", tt.vals, got, tt.want)
			}
			for i := range got {
				if math.Abs(got[i]-tt.want[i]) > eps {
					t.Errorf("AverageOfLevels(%v)[%d] = %v, want %v", tt.vals, i, got[i], tt.want[i])
				}
			}
		})
	}
}

// buildCompleteTree 构造深度为 depth 的满二叉树，用于基准测试
func buildCompleteTree(depth int) *datastructures.TreeNode {
	val := 1
	var build func(d int) *datastructures.TreeNode
	build = func(d int) *datastructures.TreeNode {
		if d == 0 {
			return nil
		}
		node := &datastructures.TreeNode{Val: val}
		val++
		node.Left = build(d - 1)
		node.Right = build(d - 1)
		return node
	}
	return build(depth)
}

func BenchmarkAverageOfLevels(b *testing.B) {
	benchmarks := []struct {
		name  string
		depth int
	}{
		{"depth=4(15节点)", 4},
		{"depth=10(1023节点)", 10},
		{"depth=14(16383节点)", 14},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			root := buildCompleteTree(bm.depth)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				AverageOfLevels(root)
			}
		})
	}
}
