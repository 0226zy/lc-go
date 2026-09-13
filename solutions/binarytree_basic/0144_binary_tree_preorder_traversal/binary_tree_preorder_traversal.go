package preordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

type TreeNode = datastructures.TreeNode

// PreorderTraversal 二叉树的前序遍历
// 给定二叉树的根节点 root，返回其节点值的前序遍历结果（根 -> 左 -> 右）。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高，最坏 O(n)
func PreorderTraversal(root *TreeNode) []int {
	result := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

// PreorderTraversalIterative 迭代法实现二叉树前序遍历（显式栈）
// 栈中先压入右子节点再压入左子节点，保证弹出顺序为「根 -> 左 -> 右」。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高，最坏 O(n)
func PreorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}
