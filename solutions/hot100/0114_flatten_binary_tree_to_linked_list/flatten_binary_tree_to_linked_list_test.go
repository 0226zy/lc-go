package flattenbinarytreetolinkedlist

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildTree 辅助函数：从切片构建二叉树，math.MinInt32 表示 nil
func buildTree(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

// toList 将展开后的链表转为切片（只走右指针）
func toList(root *datastructures.TreeNode) []int {
	var result []int
	for node := root; node != nil; node = node.Right {
		result = append(result, node.Val)
	}
	return result
}

// cloneTree 深拷贝一棵树
func cloneTree(node *datastructures.TreeNode) *datastructures.TreeNode {
	if node == nil {
		return nil
	}
	return &datastructures.TreeNode{
		Val:   node.Val,
		Left:  cloneTree(node.Left),
		Right: cloneTree(node.Right),
	}
}

// verifyFlattened 验证展开后的链表：左指针全为 nil，右指针形成链表
func verifyFlattened(root *datastructures.TreeNode) bool {
	for node := root; node != nil; node = node.Right {
		if node.Left != nil {
			return false
		}
	}
	return true
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name     string
		root     *datastructures.TreeNode
		expected []int
	}{
		{
			name:     "空树",
			root:     nil,
			expected: nil,
		},
		{
			name:     "单节点",
			root:     buildTree([]int{1}),
			expected: []int{1},
		},
		{
			name:     "官方示例1",
			root:     buildTree([]int{1, 2, 5, 3, 4, math.MinInt32, 6}),
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "只有左子树",
			root:     buildTree([]int{1, 2, math.MinInt32, 3}),
			expected: []int{1, 2, 3},
		},
		{
			name:     "只有右子树",
			root:     buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			expected: []int{1, 2, 3},
		},
		{
			name:     "完全二叉树",
			root:     buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			expected: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name:     "左斜树",
			root:     buildTree([]int{1, 2, math.MinInt32, 3, math.MinInt32, 4}),
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "右斜树",
			root:     buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3, math.MinInt32, 4}),
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Flatten(tt.root)
			got := toList(tt.root)
			if !equalSlice(got, tt.expected) {
				t.Errorf("Flatten() = %v, want %v", got, tt.expected)
			}
			if !verifyFlattened(tt.root) {
				t.Errorf("Flatten() 展开后左指针未全部置空")
			}
		})
	}
}

func TestFlattenIterative(t *testing.T) {
	tests := []struct {
		name     string
		root     *datastructures.TreeNode
		expected []int
	}{
		{
			name:     "空树",
			root:     nil,
			expected: nil,
		},
		{
			name:     "单节点",
			root:     buildTree([]int{1}),
			expected: []int{1},
		},
		{
			name:     "官方示例1",
			root:     buildTree([]int{1, 2, 5, 3, 4, math.MinInt32, 6}),
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "只有左子树",
			root:     buildTree([]int{1, 2, math.MinInt32, 3}),
			expected: []int{1, 2, 3},
		},
		{
			name:     "只有右子树",
			root:     buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3}),
			expected: []int{1, 2, 3},
		},
		{
			name:     "完全二叉树",
			root:     buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			expected: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name:     "左斜树",
			root:     buildTree([]int{1, 2, math.MinInt32, 3, math.MinInt32, 4}),
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "右斜树",
			root:     buildTree([]int{1, math.MinInt32, 2, math.MinInt32, 3, math.MinInt32, 4}),
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenIterative(tt.root)
			got := toList(tt.root)
			if !equalSlice(got, tt.expected) {
				t.Errorf("FlattenIterative() = %v, want %v", got, tt.expected)
			}
			if !verifyFlattened(tt.root) {
				t.Errorf("FlattenIterative() 展开后左指针未全部置空")
			}
		})
	}
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkFlatten(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		root := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
		b.StartTimer()
		Flatten(root)
	}
}

func BenchmarkFlattenIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		root := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
		b.StartTimer()
		FlattenIterative(root)
	}
}
