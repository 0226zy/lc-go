package convertsortedarraytobinarysearchtree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// isBST 辅助函数：验证是否为有效 BST
func isBST(root *datastructures.TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max)
}

// treeHeight 辅助函数：计算树的高度
func treeHeight(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeHeight(root.Left)
	right := treeHeight(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// isBalanced 辅助函数：验证是否为平衡二叉树
func isBalanced(root *datastructures.TreeNode) bool {
	if root == nil {
		return true
	}
	diff := treeHeight(root.Left) - treeHeight(root.Right)
	if diff < -1 || diff > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

// treeSize 辅助函数：计算树的节点数
func treeSize(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + treeSize(root.Left) + treeSize(root.Right)
}

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "官方示例1",
			nums: []int{-10, -3, 0, 5, 9},
		},
		{
			name: "官方示例2",
			nums: []int{1, 3},
		},
		{
			name: "单元素",
			nums: []int{0},
		},
		{
			name: "三个元素",
			nums: []int{-1, 0, 1},
		},
		{
			name: "五个元素",
			nums: []int{-2, -1, 0, 1, 2},
		},
		{
			name: "七个元素",
			nums: []int{-3, -2, -1, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortedArrayToBST(tt.nums)
			// 验证是有效 BST
			if !isBST(got, math.MinInt32, math.MaxInt32) {
				t.Errorf("SortedArrayToBST() 结果不是有效 BST")
			}
			// 验证是平衡树
			if !isBalanced(got) {
				t.Errorf("SortedArrayToBST() 结果不是平衡树")
			}
			// 验证节点数一致
			if treeSize(got) != len(tt.nums) {
				t.Errorf("SortedArrayToBST() 节点数 = %v, want %v", treeSize(got), len(tt.nums))
			}
		})
	}
}

func TestSortedArrayToBSTIterative(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "官方示例1",
			nums: []int{-10, -3, 0, 5, 9},
		},
		{
			name: "官方示例2",
			nums: []int{1, 3},
		},
		{
			name: "单元素",
			nums: []int{0},
		},
		{
			name: "三个元素",
			nums: []int{-1, 0, 1},
		},
		{
			name: "五个元素",
			nums: []int{-2, -1, 0, 1, 2},
		},
		{
			name: "七个元素",
			nums: []int{-3, -2, -1, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortedArrayToBSTIterative(tt.nums)
			// 验证是有效 BST
			if !isBST(got, math.MinInt32, math.MaxInt32) {
				t.Errorf("SortedArrayToBSTIterative() 结果不是有效 BST")
			}
			// 验证是平衡树
			if !isBalanced(got) {
				t.Errorf("SortedArrayToBSTIterative() 结果不是平衡树")
			}
			// 验证节点数一致
			if treeSize(got) != len(tt.nums) {
				t.Errorf("SortedArrayToBSTIterative() 节点数 = %v, want %v", treeSize(got), len(tt.nums))
			}
		})
	}
}

func BenchmarkSortedArrayToBST(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i - 500
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortedArrayToBST(nums)
	}
}

func BenchmarkSortedArrayToBSTIterative(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i - 500
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortedArrayToBSTIterative(nums)
	}
}
