package buildtree

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestBuildTree(t *testing.T) {
	nilNode := math.MinInt32
	tests := []struct {
		name      string
		inorder   []int
		postorder []int
		want      []int // 层序遍历（math.MinInt32 表示空节点）
	}{
		// LeetCode 官方示例
		{"示例1: 五节点树", []int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}, []int{3, 9, 20, nilNode, nilNode, 15, 7}},
		{"示例2: 单负数节点", []int{-1}, []int{-1}, []int{-1}},

		// 边界情况
		{"空树", nil, nil, nil},
		{"两节点只有左孩子", []int{2, 1}, []int{2, 1}, []int{1, 2}},
		{"两节点只有右孩子", []int{1, 2}, []int{2, 1}, []int{1, nilNode, 2}},
		{"左斜树", []int{3, 2, 1}, []int{3, 2, 1}, []int{1, 2, nilNode, 3}},
		{"右斜树", []int{1, 2, 3}, []int{3, 2, 1}, []int{1, nilNode, 2, nilNode, 3}},
		{"含负数与极值", []int{-3000, 0, 3000}, []int{-3000, 3000, 0}, []int{0, -3000, 3000}},
		{"七节点满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 3, 2, 5, 7, 6, 4}, []int{4, 2, 6, 1, 3, 5, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := flatten(datastructures.NewTreeFromSlice(tt.want))
			root := BuildTree(tt.inorder, tt.postorder)
			if result := flatten(root); !reflect.DeepEqual(result, want) {
				t.Errorf("BuildTree(%v, %v) = %v, want %v", tt.inorder, tt.postorder, result, want)
			}
		})
	}
}

// flatten 将树的层序遍历展平为一维切片（math.MinInt32 表示空节点），便于比较；
// 末尾多余的空节点占位会被裁掉，nil 树返回 nil 切片
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
	for len(result) > 0 && result[len(result)-1] == math.MinInt32 {
		result = result[:len(result)-1]
	}
	return result
}

func BenchmarkBuildTree(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"节点数100", 100},
		{"节点数1000", 1000},
		{"节点数3000", 3000},
	}
	for _, bm := range benchmarks {
		inorder, postorder := generateTraversalPair(bm.n)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BuildTree(inorder, postorder)
			}
		})
	}
}

// generateTraversalPair 生成一组形状随机的二叉树对应的中序/后序序列对：
// 打乱 1..n 后依次插入二叉搜索树，再遍历得到两个序列
func generateTraversalPair(n int) (inorder, postorder []int) {
	inorder = make([]int, n)
	for i := range inorder {
		inorder[i] = i + 1
	}
	randShuffle(inorder)
	var root *datastructures.TreeNode
	for _, v := range inorder {
		root = insertBST(root, v)
	}
	return inorderOf(root), postorderOf(root)
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

func inorderOf(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(inorderOf(root.Left), append([]int{root.Val}, inorderOf(root.Right)...)...)
}

func postorderOf(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(postorderOf(root.Left), append(postorderOf(root.Right), root.Val)...)
}

// randShuffle 用线性同余伪随机数打乱切片，保证基准测试可复现
func randShuffle(a []int) {
	seed := 12345
	for i := len(a) - 1; i > 0; i-- {
		seed = (seed*1103515245 + 12345) & 0x7fffffff
		j := seed % (i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
