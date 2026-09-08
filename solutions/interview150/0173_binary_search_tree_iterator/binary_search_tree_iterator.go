package binarysearchtreeiterator

import "github.com/0226zy/lc-go/pkg/datastructures"

// BSTIterator 二叉搜索树迭代器
// 按中序遍历（升序）逐个返回二叉搜索树中的值。
// Next() 均摊时间复杂度 O(1)，HasNext() 时间复杂度 O(1)，空间复杂度 O(h)（h 为树高）。
type BSTIterator struct {
	stack []*datastructures.TreeNode // 保存尚未输出的节点
}

// Constructor 初始化迭代器，把根节点所在左链路压栈
// 时间复杂度: O(h)  空间复杂度: O(h)
func Constructor(root *datastructures.TreeNode) BSTIterator {
	it := BSTIterator{}
	it.pushLeft(root)
	return it
}

// pushLeft 把 node 及其所有左孩子一路压栈
func (it *BSTIterator) pushLeft(node *datastructures.TreeNode) {
	for node != nil {
		it.stack = append(it.stack, node)
		node = node.Left
	}
}

// Next 返回中序遍历的下一个元素
// 时间复杂度: 均摊 O(1)  空间复杂度: O(1)
func (it *BSTIterator) Next() int {
	n := len(it.stack)
	top := it.stack[n-1] // 栈顶即当前最小值
	it.stack = it.stack[:n-1]
	it.pushLeft(top.Right) // 转向右子树，继续压左链
	return top.Val
}

// HasNext 判断是否还有下一个元素
// 时间复杂度: O(1)  空间复杂度: O(1)
func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}
