package buildtree

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// equalTree 递归比较两棵树的结构与节点值是否完全一致
func equalTree(a, b *datastructures.TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && equalTree(a.Left, b.Left) && equalTree(a.Right, b.Right)
}

func TestBuildTree(t *testing.T) {
	nilNode := math.MinInt32
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     []int // 层序表示，nilNode 表示空节点
	}{
		// LeetCode 官方示例
		{"示例1: 五节点树", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7},
			[]int{3, 9, 20, nilNode, nilNode, 15, 7}},
		{"示例2: 单节点负数", []int{-1}, []int{-1}, []int{-1}},

		// 边界情况
		{"空树", nil, nil, nil},
		{"单节点", []int{42}, []int{42}, []int{42}},
		{"两节点: 根只有左孩子", []int{1, 2}, []int{2, 1}, []int{1, 2}},
		{"两节点: 根只有右孩子", []int{1, 2}, []int{1, 2}, []int{1, nilNode, 2}},
		{"左链式树", []int{1, 2, 3, 4}, []int{4, 3, 2, 1},
			[]int{1, 2, nilNode, 3, nilNode, 4}},
		{"右链式树", []int{1, 2, 3, 4}, []int{1, 2, 3, 4},
			[]int{1, nilNode, 2, nilNode, 3, nilNode, 4}},
		{"含负数与极值", []int{0, -3000, 3000}, []int{-3000, 0, 3000},
			[]int{0, -3000, 3000}},
		{"七节点完全二叉树", []int{4, 2, 1, 3, 6, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7},
			[]int{4, 2, 6, 1, 3, 5, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildTree(tt.preorder, tt.inorder)
			want := datastructures.NewTreeFromSlice(tt.want)
			if !equalTree(got, want) {
				t.Errorf("BuildTree(%v, %v) 结果不符:\n得到:\n%v\n期望:\n%v",
					tt.preorder, tt.inorder, got, want)
			}
		})
	}
}

// BenchmarkBuildTree 基准测试：还原一棵 3000 节点的平衡树（题目数据规模上限）
func BenchmarkBuildTree(b *testing.B) {
	preorder, inorder := genBalancedTraversal(3000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuildTree(preorder, inorder)
	}
}

// genBalancedTraversal 生成一棵含 n 个节点的平衡二叉树的前序/中序序列
// 中序固定为 1..n，每次取区间中点为根，保证树形平衡
func genBalancedTraversal(n int) (preorder, inorder []int) {
	inorder = make([]int, n)
	for i := range inorder {
		inorder[i] = i + 1
	}
	var walk func(lo, hi int)
	walk = func(lo, hi int) {
		if lo > hi {
			return
		}
		mid := lo + (hi-lo)/2
		preorder = append(preorder, inorder[mid])
		walk(lo, mid-1)
		walk(mid+1, hi)
	}
	walk(0, n-1)
	return preorder, inorder
}
