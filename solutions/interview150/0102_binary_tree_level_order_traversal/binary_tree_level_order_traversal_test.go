package binarytreelevelordertraversal

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 三层二叉树", []int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"示例2: 单节点", []int{1}, [][]int{{1}}},
		// 边界：空树
		{"空树", nil, nil},

		// 边界：只有左子树
		{"只有左子树", []int{1, 2, math.MinInt32, 3}, [][]int{{1}, {2}, {3}}},
		// 边界：只有右子树
		{"只有右子树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, [][]int{{1}, {2}, {3}}},
		// 边界：满二叉树
		{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
		// 边界：同一父节点只有一个孩子（交错缺子）
		{"交错缺子", []int{1, 2, 3, 4, math.MinInt32, math.MinInt32, 5}, [][]int{{1}, {2, 3}, {4, 5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := LevelOrder(root)
			if tt.want == nil {
				if got != nil && len(got) != 0 {
					t.Errorf("LevelOrder(%v) = %v, want nil", tt.vals, got)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder(%v) = %v, want %v", tt.vals, got, tt.want)
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

func BenchmarkLevelOrder(b *testing.B) {
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
				LevelOrder(root)
			}
		})
	}
}
