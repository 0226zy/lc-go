package maximumbinarytree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// isSameTree 递归比较两棵树的结构与值是否完全一致（区分 nil 位置，避免层序遍历丢失结构信息）
func isSameTree(a, b *datastructures.TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *datastructures.TreeNode
	}{
		// LeetCode 官方示例
		{"示例1: nums=[3,2,1,6,0,5]",
			[]int{3, 2, 1, 6, 0, 5},
			datastructures.NewTreeFromSlice([]int{6, 3, 5, math.MinInt32, 2, 0, math.MinInt32, math.MinInt32, 1})},
		{"示例2: nums=[3,2,1]",
			[]int{3, 2, 1},
			datastructures.NewTreeFromSlice([]int{3, math.MinInt32, 2, math.MinInt32, 1})},

		// 边界：单节点
		{"单节点", []int{5}, datastructures.NewTreeFromSlice([]int{5})},
		// 边界：升序数组，退化为全左链
		{"升序数组-全左链", []int{1, 2, 3},
			datastructures.NewTreeFromSlice([]int{3, 2, math.MinInt32, 1})},
		// 边界：降序数组，退化为全右链
		{"降序数组-全右链", []int{3, 2, 1},
			datastructures.NewTreeFromSlice([]int{3, math.MinInt32, 2, math.MinInt32, 1})},
		// 边界：含负数
		{"含负数", []int{-3, -1, -2},
			datastructures.NewTreeFromSlice([]int{-1, -3, -2})},
		// 边界：最大值在中间
		{"最大值在中间", []int{1, 5, 3},
			datastructures.NewTreeFromSlice([]int{5, 1, 3})},
		// 边界：最大值在左端
		{"最大值在左端", []int{5, 1, 3},
			datastructures.NewTreeFromSlice([]int{5, math.MinInt32, 3, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConstructMaximumBinaryTree(tt.nums)
			if !isSameTree(got, tt.want) {
				t.Errorf("ConstructMaximumBinaryTree(%v) = %v, want %v", tt.nums, got.LevelOrder(), tt.want.LevelOrder())
			}
		})
	}
}

// TestConstructMaximumBinaryTreeEmpty 空数组应返回 nil
func TestConstructMaximumBinaryTreeEmpty(t *testing.T) {
	if got := ConstructMaximumBinaryTree([]int{}); got != nil {
		t.Errorf("ConstructMaximumBinaryTree([]) = %v, want nil", got.LevelOrder())
	}
	if got := ConstructMaximumBinaryTree(nil); got != nil {
		t.Errorf("ConstructMaximumBinaryTree(nil) = %v, want nil", got.LevelOrder())
	}
}

func BenchmarkConstructMaximumBinaryTree(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		// 升序数组是最坏情况：每次扫描区间长度递减，总比较次数 O(n²)
		{"升序数组100", makeAscending(100)},
		{"升序数组500", makeAscending(500)},
		{"升序数组1000", makeAscending(1000)},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConstructMaximumBinaryTree(bm.nums)
			}
		})
	}
}

// makeAscending 生成 1..n 的升序数组，用于构造最坏情况
func makeAscending(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	return nums
}
