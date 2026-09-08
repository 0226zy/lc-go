package constructbinarytreefrompreorderandinordertraversal

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestBuildTree(t *testing.T) {
	nilNode := math.MinInt32
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     []int
	}{
		// LeetCode 官方示例
		{"示例1: 五节点树", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []int{3, 9, 20, nilNode, nilNode, 15, 7}},
		{"示例2: 单节点", []int{-1}, []int{-1}, []int{-1}},

		// 边界情况
		{"两节点", []int{1, 2}, []int{2, 1}, []int{1, 2}},
		{"左斜树", []int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2, nilNode, 3}},
		{"右斜树", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, nilNode, 2, nilNode, 3}},
		{"负数与极值", []int{0, -3000, 3000}, []int{-3000, 0, 3000}, []int{0, -3000, 3000}},
		{"七节点树", []int{4, 2, 1, 3, 6, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}, []int{4, 2, 6, 1, 3, 5, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flatten(datastructures.NewTreeFromSlice(tt.want))
			root := BuildTree(tt.preorder, tt.inorder)
			if result := flatten(root); !reflect.DeepEqual(result, got) {
				t.Errorf("BuildTree(%v, %v) = %v, want %v", tt.preorder, tt.inorder, result, got)
			}
		})
	}
}

// flatten 展平层序遍历结果为一维切片（含 nil 占位），便于比较
func flatten(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, child := range []*datastructures.TreeNode{node.Left, node.Right} {
			if child != nil {
				result = append(result, child.Val)
				queue = append(queue, child)
			} else {
				result = append(result, math.MinInt32)
			}
		}
	}
	// 去掉末尾多余的 nil 占位
	for len(result) > 0 && result[len(result)-1] == math.MinInt32 {
		result = result[:len(result)-1]
	}
	return result
}

func BenchmarkBuildTree(b *testing.B) {
	preorder, inorder := generateTraversalPair(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuildTree(preorder, inorder)
	}
}

// generateTraversalPair 生成随机搜索树形状的先序/中序序列对
func generateTraversalPair(n int) (preorder, inorder []int) {
	inorder = make([]int, n)
	for i := range inorder {
		inorder[i] = i + 1
	}
	// 随机打乱中序，作为插入顺序，使树形随机
	randShuffle(inorder)
	// 以 inorder 的顺序依次插入 1..n，构造随机 BST，再遍历出先序/中序
	var root *datastructures.TreeNode
	for _, v := range inorder {
		root = insertBST(root, v)
	}
	return preorderOf(root), inorderOf(root)
}

func insertBST(root *datastructures.TreeNode, val int) *datastructures.TreeNode {
	if root == nil {
		return &datastructures.TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertBST(root.Left, val)
	} else {
		root.Right = insertBST(root.Right, val)
	}
	return root
}

func preorderOf(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(append([]int{root.Val}, preorderOf(root.Left)...), preorderOf(root.Right)...)
}

func inorderOf(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(inorderOf(root.Left), append([]int{root.Val}, inorderOf(root.Right)...)...)
}

// randShuffle 用简单的线性同余伪随机打乱切片（避免引入 math/rand 也可，但标准库允许）
func randShuffle(a []int) {
	seed := 12345
	for i := len(a) - 1; i > 0; i-- {
		seed = (seed*1103515245 + 12345) & 0x7fffffff
		j := seed % (i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
