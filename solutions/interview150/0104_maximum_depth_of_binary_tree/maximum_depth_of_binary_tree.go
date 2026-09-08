package maximumdepthofbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// MaxDepth 二叉树的最大深度
// 给定一个二叉树 root，返回其最大深度（从根节点到最远叶子节点的最长路径上的节点总数）。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func MaxDepth(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
