package invertbinarytree

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

// treeToSlice 辅助函数：将二叉树转为层序切片，便于比较
func treeToSlice(root *datastructures.TreeNode) []int {
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

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want []int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: []int{1},
		},
		{
			name: "官方示例1",
			root: buildTree([]int{4, 2, 7, 1, 3, 6, 9}),
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "官方示例2",
			root: buildTree([]int{2, 1, 3}),
			want: []int{2, 3, 1},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{1, 3, 2, 7, 6, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := treeToSlice(InvertTree(tt.root))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvertTreeBFS(t *testing.T) {
	tests := []struct {
		name string
		root *datastructures.TreeNode
		want []int
	}{
		{
			name: "空树",
			root: nil,
			want: nil,
		},
		{
			name: "单节点",
			root: buildTree([]int{1}),
			want: []int{1},
		},
		{
			name: "官方示例1",
			root: buildTree([]int{4, 2, 7, 1, 3, 6, 9}),
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "官方示例2",
			root: buildTree([]int{2, 1, 3}),
			want: []int{2, 3, 1},
		},
		{
			name: "左斜树",
			root: buildTree([]int{1, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "右斜树",
			root: buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "完全二叉树",
			root: buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{1, 3, 2, 7, 6, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := treeToSlice(InvertTreeBFS(tt.root))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertTreeBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInvertTree(b *testing.B) {
	root := buildTree([]int{4, 2, 7, 1, 3, 6, 9})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InvertTree(root)
	}
}

func BenchmarkInvertTreeBFS(b *testing.B) {
	root := buildTree([]int{4, 2, 7, 1, 3, 6, 9})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InvertTreeBFS(root)
	}
}
