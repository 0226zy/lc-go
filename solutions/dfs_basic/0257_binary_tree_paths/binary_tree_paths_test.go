package binarytreepaths

import (
	"math"
	"sort"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// normalize 对字符串切片排序后返回，用于结果的无序比较
func normalize(paths []string) []string {
	cp := make([]string, len(paths))
	copy(cp, paths)
	sort.Strings(cp)
	return cp
}

// equalPaths 无序比较两个字符串切片（nil 与空切片视为相等）
func equalPaths(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sa, sb := normalize(a), normalize(b)
	for i := range sa {
		if sa[i] != sb[i] {
			return false
		}
	}
	return true
}

func TestBinaryTreePaths(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []string
	}{
		// LeetCode 官方示例
		{"示例1: 三条路径", []int{1, 2, 3, math.MinInt32, 5}, []string{"1->2->5", "1->3"}},
		{"示例2: 单节点", []int{1}, []string{"1"}},
		// 边界：空树
		{"空树", []int{}, []string{}},
		// 边界：只有左孩子的链
		{"左斜树", []int{1, 2, math.MinInt32, 3}, []string{"1->2->3"}},
		// 边界：只有右孩子的链
		{"右斜树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, []string{"1->2->3"}},
		// 边界：左右子树各一条路径
		{"左右各一叶", []int{1, 2, 3}, []string{"1->2", "1->3"}},
		// 边界：负数节点值
		{"负数节点", []int{-1, -2, 3}, []string{"-1->-2", "-1->3"}},
		// 边界：完全二叉树，四条路径
		{"完全二叉树", []int{1, 2, 3, 4, 5, 6, 7}, []string{"1->2->4", "1->2->5", "1->3->6", "1->3->7"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := BinaryTreePaths(root)
			if !equalPaths(got, tt.want) {
				t.Errorf("BinaryTreePaths(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkBinaryTreePaths(b *testing.B) {
	// 构建一棵约 4095 个节点的完全二叉树，共 2048 条根到叶路径
	var vals []int
	for size := 1; len(vals)+size <= 4095; size *= 2 {
		for i := 0; i < size; i++ {
			vals = append(vals, i+1)
		}
	}
	root := datastructures.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinaryTreePaths(root)
	}
}
