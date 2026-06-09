package kthsmallestelementinabst

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// KthSmallest 二叉搜索树中第 K 小的元素（中序遍历递归）
// BST 中序遍历结果按升序排列，第 k 个即为答案。
// 时间复杂度: O(h + k)  空间复杂度: O(h)
func KthSmallest(root *datastructures.TreeNode, k int) int {
	var result int
	var count int
	var inorder func(node *datastructures.TreeNode)
	inorder = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}
		if count < k {
			inorder(node.Right)
		}
	}
	inorder(root)
	return result
}

// KthSmallestIterative 二叉搜索树中第 K 小的元素（中序遍历迭代）
// 使用栈模拟中序遍历，找到第 k 个元素后立即停止。
// 时间复杂度: O(h + k)  空间复杂度: O(h)
func KthSmallestIterative(root *datastructures.TreeNode, k int) int {
	stack := []*datastructures.TreeNode{}
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return curr.Val
		}
		curr = curr.Right
	}
	return -1
}
