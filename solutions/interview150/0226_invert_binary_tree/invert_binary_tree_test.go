package invertbinarytree

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestInvertTree(t *testing.T) {
	nilNode := math.MinInt32
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: 三层满树", []int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{"示例2: 三节点树", []int{2, 1, 3}, []int{2, 3, 1}},
		{"示例3: 空树", nil, nil},

		// 边界情况
		{"单节点", []int{1}, []int{1}},
		{"左斜树", []int{1, 2, nilNode, 3}, []int{1, nilNode, 2, nilNode, 3}},
		{"负值节点", []int{-1, -2, -3}, []int{-1, -3, -2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := levelOrder(InvertTree(root))
			want := levelOrder(datastructures.NewTreeFromSlice(tt.want))
			if !reflect.DeepEqual(got, want) {
				t.Errorf("InvertTree(%v) = %v, want %v", tt.vals, got, want)
			}
		})
	}
}

// levelOrder 展平层序遍历结果为一维切片，便于比较
func levelOrder(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

func BenchmarkInvertTree(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"节点数7", []int{4, 2, 7, 1, 3, 6, 9}},
		{"节点数63", buildCompleteTreeVals(6)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				root := datastructures.NewTreeFromSlice(bm.vals)
				InvertTree(root)
			}
		})
	}
}

// buildCompleteTreeVals 生成深度为 depth 的满二叉树的层序序列
func buildCompleteTreeVals(depth int) []int {
	n := 1 << depth
	vals := make([]int, n-1)
	for i := range vals {
		vals[i] = i + 1
	}
	return vals
}
