package postordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

type TreeNode = datastructures.TreeNode

// PostorderTraversal 二叉树的后序遍历
// 给定二叉树的根节点 root，返回其节点值的后序遍历序列（左子树 -> 右子树 -> 根节点）。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高，最坏 O(n)
func PostorderTraversal(root *TreeNode) []int {
	result := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}
	dfs(root)
	return result
}

// PostorderTraversalIterative 迭代法实现二叉树后序遍历（显式栈 + 结果逆置）
// 先按「根 -> 右 -> 左」的顺序入栈遍历，最后将结果反转，即得到「左 -> 右 -> 根」。
// 时间复杂度: O(n)  空间复杂度: O(h)，h 为树高，最坏 O(n)
func PostorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// 反转结果：「根 -> 右 -> 左」逆置后即为「左 -> 右 -> 根」
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
