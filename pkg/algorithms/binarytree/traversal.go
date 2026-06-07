package binarytree

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// TreeTraversal 封装二叉树常用遍历算法
type TreeTraversal struct{}

// NewTreeTraversal 创建一个新的 TreeTraversal 实例
func NewTreeTraversal() *TreeTraversal {
	return &TreeTraversal{}
}

// PreorderRecursive 前序遍历（递归）
// 遍历顺序: 根 -> 左 -> 右
func (tt *TreeTraversal) PreorderRecursive(root *datastructures.TreeNode) []int {
	// TODO: 手写实现
	panic("not implemented")
}

// PreorderIterative 前序遍历（非递归）
// 使用栈模拟递归过程
func (tt *TreeTraversal) PreorderIterative(root *datastructures.TreeNode) []int {
	// TODO: 手写实现
	panic("not implemented")
}

// InorderRecursive 中序遍历（递归）
// 遍历顺序: 左 -> 根 -> 右
func (tt *TreeTraversal) InorderRecursive(root *datastructures.TreeNode) []int {
	// TODO: 手写实现
	panic("not implemented")
}

// InorderIterative 中序遍历（非递归）
// 使用栈模拟递归过程
func (tt *TreeTraversal) InorderIterative(root *datastructures.TreeNode) []int {
	// TODO: 手写实现
	panic("not implemented")
}

// PostorderRecursive 后序遍历（递归）
// 遍历顺序: 左 -> 右 -> 根
func (tt *TreeTraversal) PostorderRecursive(root *datastructures.TreeNode) []int {
	// TODO: 手写实现
	panic("not implemented")
}

// PostorderIterative 后序遍历（非递归）
// 使用栈模拟递归过程
func (tt *TreeTraversal) PostorderIterative(root *datastructures.TreeNode) []int {
	// TODO: 手写实现
	panic("not implemented")
}

// LevelOrder 层序遍历
// 从上到下、从左到右依次遍历每个节点
func (tt *TreeTraversal) LevelOrder(root *datastructures.TreeNode) []int {
	// TODO: 手写实现
	panic("not implemented")
}

// LevelOrderByLevel 层序遍历，按层返回结果
// 每一层的节点值放在一个单独的切片中
func (tt *TreeTraversal) LevelOrderByLevel(root *datastructures.TreeNode) [][]int {
	// TODO: 手写实现
	panic("not implemented")
}

// LevelOrderZigzag 层序遍历之字形（锯齿形）
// 奇数层从左到右，偶数层从右到左
func (tt *TreeTraversal) LevelOrderZigzag(root *datastructures.TreeNode) [][]int {
	// TODO: 手写实现
	panic("not implemented")
}
